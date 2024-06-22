package database

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(host string, port string, password string, db int) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Redis{
		Client: client,
	}, nil
}
