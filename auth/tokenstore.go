package auth

import (
	"time"

	"github.com/TerrexTech/go-apigateway/gql/entity/auth/model"
	"github.com/TerrexTech/uuuid"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

// TokenStoreI is a key-value token-storage.
type TokenStoreI interface {
	Set(token *model.RefreshToken) error
	Get(uid uuuid.UUID) (string, error)
}

// Redis represents a Redis-Client instance.
// This implements the TokenStoreI.
type Redis struct {
	client *redis.Client
}

// NewRedis creates a new Redis-Client using provided options.
func NewRedis(opts *redis.Options) (*Redis, error) {
	client := redis.NewClient(opts)
	if client == nil {
		return nil, errors.New("Redis Error: Unable to create client")
	}
	return &Redis{
		client: client,
	}, nil
}

// Set sets the provided token with "sub" as key.
func (r *Redis) Set(token *model.RefreshToken) error {
	exp := time.Until(token.Exp)
	status := r.client.Set(token.Sub.String(), token.Token, exp)
	err := status.Err()
	if err != nil {
		err = errors.Wrap(err, "Redis Error in Set")
		return err
	}
	return nil
}

// Get retrieves token from Redis using provided uuid as key.
func (r *Redis) Get(uid uuuid.UUID) (string, error) {
	val, err := r.client.Get(uid.String()).Result()
	if err != nil {
		err = errors.Wrap(err, "Redis Error in Get")
		return "", err
	}
	return val, nil
}
