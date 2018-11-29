package main

import (
	"log"
	"os"
	"strconv"

	"github.com/TerrexTech/go-commonutils/commonutil"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/pkg/errors"
)

func loadMongoConfig() (*mongo.Collection, error) {
	hosts := *commonutil.ParseHosts(
		os.Getenv("MONGO_HOSTS"),
	)

	database := os.Getenv("MONGO_DATABASE")
	aggCollection := os.Getenv("MONGO_INV_COLLECTION")
	username := os.Getenv("MONGO_USERNAME")
	password := os.Getenv("MONGO_PASSWORD")
	connTimeoutStr := os.Getenv("MONGO_CONNECTION_TIMEOUT_MS")
	connTimeout, err := strconv.Atoi(connTimeoutStr)
	if err != nil {
		err = errors.Wrap(err, "Error converting MONGO_CONNECTION_TIMEOUT_MS to integer")
		log.Println(err)
		log.Println("A defalt value of 3000 will be used for MONGO_CONNECTION_TIMEOUT_MS")
		connTimeout = 3000
	}

	mongoConfig := mongo.ClientConfig{
		Hosts:               hosts,
		Username:            username,
		Password:            password,
		TimeoutMilliseconds: uint32(connTimeout),
	}

	// MongoDB Client
	client, err := mongo.NewClient(mongoConfig)
	if err != nil {
		err = errors.Wrap(err, "Error creating MongoClient")
		log.Fatalln(err)
	}

	resTimeoutStr := os.Getenv("MONGO_CONNECTION_TIMEOUT_MS")
	resTimeout, err := strconv.Atoi(resTimeoutStr)
	if err != nil {
		err = errors.Wrap(err, "Error converting MONGO_RESOURCE_TIMEOUT_MS to integer")
		log.Println(err)
		log.Println("A defalt value of 5000 will be used for MONGO_RESOURCE_TIMEOUT_MS")
		connTimeout = 5000
	}
	conn := &mongo.ConnectionConfig{
		Client:  client,
		Timeout: uint32(resTimeout),
	}

	aggMongoCollection, err := createMongoCollection(conn, database, aggCollection)
	if err != nil {
		err = errors.Wrap(err, "Error creating MongoCollection")
		return nil, err
	}

	return aggMongoCollection, nil
}

func createMongoCollection(
	conn *mongo.ConnectionConfig, db string, coll string,
) (*mongo.Collection, error) {
	// Index Configuration
	indexConfigs := []mongo.IndexConfig{
		mongo.IndexConfig{
			ColumnConfig: []mongo.IndexColumnConfig{
				mongo.IndexColumnConfig{
					Name: "itemID",
				},
			},
			IsUnique: true,
			Name:     "itemID_index",
		},
		mongo.IndexConfig{
			ColumnConfig: []mongo.IndexColumnConfig{
				mongo.IndexColumnConfig{
					Name:        "soldWeight",
					IsDescOrder: true,
				},
			},
			Name: "soldWeight_Index",
		},
	}

	log.Println(db)

	// Create New Collection
	c := &mongo.Collection{
		Connection:   conn,
		Database:     db,
		Name:         coll,
		SchemaStruct: &Inventory{},
		Indexes:      indexConfigs,
	}
	collection, err := mongo.EnsureCollection(c)
	if err != nil {
		err = errors.Wrap(err, "Error creating MongoCollection")
		return nil, err
	}
	return collection, nil
}

// createClient creates a MongoDB-Client.
func CreateClient(host []string, username string, password string) (*mongo.Client, error) {
	// Would ideally set these config-params as environment vars
	config := mongo.ClientConfig{
		Hosts:               host,
		Username:            username,
		Password:            password,
		TimeoutMilliseconds: 3000,
	}

	// ====> MongoDB Client
	client, err := mongo.NewClient(config)
	// Let the parent functions handle error, always -.-
	// (Even though in these examples, we won't, for simplicity)
	return client, err
}

// createCollection demonstrates creating the collection and the associated database.
func CreateCollection(client *mongo.Client,
	collName string,
	schema interface{},
	columnIndexName string,
	columnConfigName string) (*mongo.Collection, error) {
	// ====> Collection Configuration
	conn := &mongo.ConnectionConfig{
		Client:  client,
		Timeout: 5000,
	}
	// Index Configuration
	indexConfigs := []mongo.IndexConfig{
		mongo.IndexConfig{
			ColumnConfig: []mongo.IndexColumnConfig{
				mongo.IndexColumnConfig{
					Name: "itemID",
				},
				mongo.IndexColumnConfig{
					Name: columnIndexName,
				},
			},
			Name: columnConfigName,
		},
	}
	// ====> Create New Collection
	c := &mongo.Collection{
		Connection:   conn,
		Name:         collName,
		Database:     "rns_projections",
		SchemaStruct: schema,
		Indexes:      indexConfigs,
	}
	return mongo.EnsureCollection(c)
}
