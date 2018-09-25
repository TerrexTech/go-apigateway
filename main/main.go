package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/schema"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/go-redis/redis"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var (
	// Schema represents a GraphQL schema
	Schema       graphql.Schema
	rootObject   map[string]interface{}
	kafkaAdapter *kafka.Adapter
)

func initService() {
	// Load environment-file.
	// Env vars will be read directly from environment if this file fails loading
	err := godotenv.Load()
	if err != nil {
		err = errors.Wrap(err,
			".env file not found, env-vars will be read as set in environment",
		)
		log.Println(err)
	}

	missingVar, err := commonutil.ValidateEnv(
		"KAFKA_BROKERS",
		"KAFKA_CONSUMER_TOPIC_LOGIN",
		"KAFKA_PRODUCER_TOPIC_LOGIN",
		"KAFKA_CONSUMER_GROUP_REGISTER",
		"KAFKA_CONSUMER_TOPIC_REGISTER",
		"KAFKA_PRODUCER_TOPIC_REGISTER",

		"MONGO_HOSTS",
		"MONGO_DATABASE",
		"MONGO_COLLECTION",
		"MONGO_TIMEOUT",

		"REDIS_HOST",
		"REDIS_DB",
	)
	if err != nil {
		log.Fatalf(
			"Error: Environment variable %s is required but was not found", missingVar,
		)
	}

	// Redis Setup
	redisHost := os.Getenv("REDIS_HOST")
	redisDBStr := os.Getenv("REDIS_DB")
	redisDB, err := strconv.Atoi(redisDBStr)
	if err != nil {
		err = errors.Wrap(err, "Error converting REDIS_DB value to int")
		log.Fatalln(err)
	}
	tokenStore, err := auth.NewRedis(&redis.Options{
		Addr: redisHost,
		DB:   redisDB,
	})
	if err != nil {
		err = errors.Wrap(err, "Error creating Redis client")
		log.Println(err)
		return
	}

	// Kafka Setup
	brokers := os.Getenv("KAFKA_BROKERS")
	kafkaAdapter := &kafka.Adapter{
		Brokers: *commonutil.ParseHosts(brokers),
	}

	// Mongo Setup
	hosts := os.Getenv("MONGO_HOSTS")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	database := os.Getenv("MONGO_DATABASE")
	collection := os.Getenv("MONGO_COLLECTION")

	timeoutMilliStr := os.Getenv("MONGO_TIMEOUT")
	parsedTimeoutMilli, err := strconv.Atoi(timeoutMilliStr)
	if err != nil {
		err = errors.Wrap(err, "Error converting Timeout value to int32")
		log.Println(err)
		log.Println("MONGO_TIMEOUT value will be set to 3000 as default value")
		parsedTimeoutMilli = 3000
	}
	timeoutMilli := uint32(parsedTimeoutMilli)

	config := auth.DBIConfig{
		Hosts:               *commonutil.ParseHosts(hosts),
		Username:            username,
		Password:            password,
		TimeoutMilliseconds: timeoutMilli,
		Database:            database,
		Collection:          collection,
	}
	db, err := auth.EnsureAuthDB(config)
	if err != nil {
		err = errors.Wrap(err, "Error connecting to Auth-DB")
		log.Println(err)
		return
	}

	rootObject = map[string]interface{}{
		"kafkaAdapter": kafkaAdapter,
		"tokenStore":   tokenStore,
		"db":           db,
	}

	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.RootQuery,
		Mutation: schema.RootMutation,
	})
	if err != nil {
		log.Fatalf("Error creating GraphQL Schema: %v", err)
	}
	Schema = s
}

func main() {
	initService()
	port := fmt.Sprintf(":%d", 8081)
	http.HandleFunc("/api", graphqlHandler)
	err := http.ListenAndServe(port, nil)

	err = errors.Wrap(err, "Error creating server")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Listening on port " + port)
}
