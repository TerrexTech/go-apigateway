package auth

import (
	"github.com/TerrexTech/go-apigateway/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
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
	Register(user *model.User) (*model.User, error)
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

// Register inserts the provided user into database
func (d *DB) Register(user *model.User) (*model.User, error) {
	authUser := &model.User{
		ID:        objectid.New(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Username:  user.Username,
		Password:  user.Password,
	}
	uid, err := uuuid.NewV4()
	if err != nil {
		err = errors.Wrap(err, "Registration: Error generating UUID")
		return nil, err
	}
	authUser.UUID = uid

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		err = errors.Wrap(err, "Registration: Error creating Hash for password")
		return nil, err
	}
	authUser.Password = string(hashedPass)

	_, err = d.collection.InsertOne(authUser)
	if err != nil {
		err = errors.Wrap(err, "Registration: Error inserting user into Database")
		return nil, err
	}
	// Don't send hashed-password to any other service
	authUser.Password = ""
	return authUser, nil
}

// Collection returns the currrent MongoDB collection being used for user-auth operations.
func (d *DB) Collection() *mongo.Collection {
	return d.collection
}
