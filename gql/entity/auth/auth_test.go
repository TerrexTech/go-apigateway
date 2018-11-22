package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/TerrexTech/go-apigateway/auth"
	"github.com/TerrexTech/go-apigateway/gql/schema"
	"github.com/TerrexTech/go-apigateway/util"
	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/go-redis/redis"
	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth gateway Suite")
}

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

	// The GraphQL context
	rootObject = map[string]interface{}{
		"kafkaFactory": kafkaFactory,
		"tokenStore":   tokenStore,
		"appContext":   ctx,
		"errGroup":     g,
		"closeChan":    (chan<- struct{})(closeChan),
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

var _ = Describe("Auth gateway test", func() {
	var ()

	BeforeEach(func() {
		initServices()
		port := fmt.Sprintf(":%d", 8081)
		http.HandleFunc("/api", graphqlHandler)

		log.Println("Listening on port " + port)
	})

	It("Should register user", func() {
		registerInfo := `mutation{
			authRegister(
			  userName: "test",
			  password: "test",
			  firstName: "test",
			  lastName: "test",
			  email: "testcom",
			  role: "employee"
			)
			{
			  accessToken,
			  refreshToken
			}
		  }`
		resp, err := http.Post("http://localhost:8081/api", "application/json", bytes.NewBuffer([]byte(registerInfo)))
		Expect(err).ToNot(HaveOccurred())
		log.Println(resp.Body)
		body, err := ioutil.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())
	})
})
