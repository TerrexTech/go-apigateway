package auth

import (
	"time"

	"github.com/TerrexTech/go-apigateway/model"
	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type TokenStoreI interface {
	Set(token *model.RefreshToken) error
	Get(uid uuid.UUID) (string, error)
}

type Redis struct {
	client *redis.Client
}

//RedisClient connects to Redis server to store tokens
func NewRedis(opts *redis.Options) (*Redis, error) {
	client := redis.NewClient(opts)
	if client == nil {
		return nil, errors.New("Redis Error: Unable to create client")
	}
	return &Redis{
		client: client,
	}, nil
}

//SetToken sets the token for redis client
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

//GetToken retrieves token from redis db
func (r *Redis) Get(uid uuid.UUID) (string, error) {
	val, err := r.client.Get(uid.String()).Result()
	if err != nil {
		err = errors.Wrap(err, "Redis Error in Get")
		return "", err
	}
	return val, nil
}
