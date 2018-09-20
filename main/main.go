package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
var rootObject map[string]interface{}
var kafkaAdapter *kafka.Adapter

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

	tokenStore, err := auth.NewRedis(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	if err != nil {
		err = errors.Wrap(err, "Error creating Redis client")
		log.Println(err)
		return
	}

	brokers := os.Getenv("KAFKA_BROKERS")
	kafkaAdapter := &kafka.Adapter{
		Brokers: *commonutil.ParseHosts(brokers),
	}
	rootObject = map[string]interface{}{
		"kafkaAdapter": kafkaAdapter,
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
	missingVar, err := commonutil.ValidateEnv(
		"EVENT_PRODUCER_TOPIC",
		"KAFKA_BROKERS",
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
