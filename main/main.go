package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/schema"
	"github.com/TerrexTech/go-apigateway/util"
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
	kafkaFactory *util.KafkaFactory
)

func initServices() {
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
		"KAFKA_CONSUMER_GROUP_LOGIN",
		"KAFKA_CONSUMER_TOPIC_LOGIN",
		"KAFKA_PRODUCER_TOPIC_LOGIN",
		"KAFKA_CONSUMER_TOPIC_REGISTER",
		"KAFKA_PRODUCER_TOPIC_REGISTER",

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
	kafkaFactory := &util.KafkaFactory{
		Brokers: *commonutil.ParseHosts(brokers),
	}

	// The GraphQL context
	rootObject = map[string]interface{}{
		"kafkaFactory": kafkaFactory,
		"tokenStore":   tokenStore,
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
	initServices()
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
