package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/pkg/errors"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/schema"
	"github.com/TerrexTech/go-apigateway/util"
	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/go-redis/redis"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
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

		"KAFKA_PRODUCER_EVENT_TOPIC",

		"KAFKA_CONSUMER_TOPIC_USERAUTH",
		"KAFKA_CONSUMER_TOPIC_INVENTORY",

		"REDIS_HOST",
		"REDIS_DB",
		"MONGO_DEVICE_COLLECTION",
		"MONGO_METRIC_COLLECTION",
		"MONGO_DEVICE_COLLECTION",
		"MONGO_USERNAME",
		"MONGO_PASSWORD",
		"MONGO_HOSTS",
	)
	if err != nil {
		log.Fatalf(
			"Error: Environment variable %s is required but was not found", missingVar,
		)
	}

	//Mongo setup
	hosts := *commonutil.ParseHosts(
		os.Getenv("MONGO_HOSTS"),
	)

	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	metricCollection := os.Getenv("MONGO_METRIC_COLLECTION")

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
	ctx, cancel := context.WithCancel(context.Background())
	closeChan := make(chan struct{})
	g, ctx := errgroup.WithContext(ctx)

	go func() {
		<-closeChan
		cancel()
	}()

	brokers := os.Getenv("KAFKA_BROKERS")
	eventProdTopic := os.Getenv("KAFKA_PRODUCER_EVENT_TOPIC")
	kafkaFactory, err := util.NewKafkaFactory(
		ctx,
		g,
		*commonutil.ParseHosts(brokers),
		eventProdTopic,
	)
	if err != nil {
		err = errors.Wrap(err, "Error creating EventProducer")
		log.Fatalln(err)
	}

	mongoColl, err := loadMongoConfig()
	if err != nil {
		err = errors.Wrap(err, "Unable to load mongocollection")
		log.Fatalln(err)
	}

	client, err := CreateClient(hosts, username, password)
	if err != nil {
		err = errors.Wrap(err, "Unable to create mongo client")
		log.Fatalln(err)
	}

	//Metric collection
	metColl, err := CreateCollection(client, metricCollection, &Metric{}, "metricID", "itemID_metricID_index")
	if err != nil {
		err = errors.Wrap(err, "Unable to create metric collection")
		log.Fatalln(err)
	}

	// The GraphQL context
	rootObject = map[string]interface{}{
		"kafkaFactory":  kafkaFactory,
		"tokenStore":    tokenStore,
		"appContext":    ctx,
		"errGroup":      g,
		"closeChan":     (chan<- struct{})(closeChan),
		"inventoryColl": mongoColl,
		"metricColl":    metColl,
	}

	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.RootQuery,
		Mutation: schema.RootMutation,
	})
	if err != nil {
		err = errors.Wrap(err, "Error creating GraphQL Schema")
		log.Fatalln(err)
	}
}

// graphqlHandler handles GraphQL requests.
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization",
		)
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = errors.Wrap(err, "Error reading request body")
		log.Println(err)
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:        Schema,
		RequestString: string(reqBytes),
		RootObject:    rootObject,
	})
	if len(result.Errors) > 0 {
		log.Printf("GQL Error: %v", result.Errors)
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		err = errors.Wrap(err, "Error writing response")
		log.Println(err)
	}
}
