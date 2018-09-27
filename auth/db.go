package auth

import (
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/pkg/errors"
)

// DBIConfig is the configuration for the authDB.
type DBIConfig struct {
	Hosts               []string
	Username            string
	Password            string
	TimeoutMilliseconds uint32
	Database            string
	Collection          string
}

// DBI is the Database-interface for authentication.
// This fetches/writes data to/from database for auth-actions such as
// login, registeration etc.
type DBI interface {
	Collection() *mongo.Collection
}

// DB is the implementation for dbI.
// DBI is the Database-interface for authentication.
// It fetches/writes data to/from database for auth-actions such as
// login, registeration etc.
type DB struct {
	collection *mongo.Collection
}

// EnsureAuthDB exists ensures that the required Database and Collection exists before
// auth-operations can be done on them. It creates Database/Collection if they don't exist.
func EnsureAuthDB(dbConfig DBIConfig) (*DB, error) {
	config := mongo.ClientConfig{
		Hosts:               dbConfig.Hosts,
		Username:            dbConfig.Username,
		Password:            dbConfig.Password,
		TimeoutMilliseconds: dbConfig.TimeoutMilliseconds,
	}

	client, err := mongo.NewClient(config)
	if err != nil {
		err = errors.Wrap(err, "Error creating DB-client")
		return nil, err
	}

	conn := &mongo.ConnectionConfig{
		Client:  client,
		Timeout: 5000,
	}

	indexConfigs := []mongo.IndexConfig{
		mongo.IndexConfig{
			ColumnConfig: []mongo.IndexColumnConfig{
				mongo.IndexColumnConfig{
					Name: "username",
				},
			},
			IsUnique: true,
			Name:     "username_index",
		},
	}

	// ====> Create New Collection
	collConfig := &mongo.Collection{
		Connection:   conn,
		Database:     dbConfig.Database,
		Name:         dbConfig.Collection,
		SchemaStruct: &model.User{},
		Indexes:      indexConfigs,
	}
	c, err := mongo.EnsureCollection(collConfig)
	if err != nil {
		err = errors.Wrap(err, "Error creating DB-client")
		return nil, err
	}
	return &DB{
		collection: c,
	}, nil
}

// Collection returns the currrent MongoDB collection being used for user-auth operations.
func (d *DB) Collection() *mongo.Collection {
	return d.collection
}
