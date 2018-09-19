package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/schema"
	"github.com/TerrexTech/go-apigateway/kafka"
	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/go-redis/redis"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

var Schema graphql.Schema
var kafkaIO *kafka.IO
var tokenStore *auth.Redis

func init() {
	// Load environment-file.
	// Env vars will be read directly from environment if this file fails loading
	err := godotenv.Load()
	if err != nil {
		err = errors.Wrap(err,
			".env file not found, env-vars will be read as set in environment",
		)
		log.Println(err)
	}

	tokenStore, err = auth.NewRedis(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err != nil {
		err = errors.Wrap(err, "Error creating Redis client")
		log.Println(err)
		return
	}
	kafkaIO, err = initKafkaIO()
	if err != nil {
		err = errors.Wrap(err, "Failed Initializing KafkaIO")
		log.Fatalln(err)
	}

	s, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.RootQuery,
	})
	if err != nil {
		log.Fatalf("Error creating GraphQL Schema: %v", err)
	}
	Schema = s
}

func main() {
	missingVar, err := commonutil.ValidateEnv(
		"KAFKA_BROKERS",
		"KAFKA_CONSUMER_GROUP_LOGIN",
		"KAFKA_CONSUMER_TOPIC_LOGIN",
		"KAFKA_PRODUCER_TOPIC_LOGIN",
	)
	if err != nil {
		log.Fatalf(
			"Error: Environment variable %s is required but was not found", missingVar,
		)
	}

	port := fmt.Sprintf(":%d", 8081)
	http.HandleFunc("/api", graphqlHandler)
	err = http.ListenAndServe(port, nil)

	err = errors.Wrap(err, "Error creating server")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Listening on port " + port)
}
