package configs

import (
	

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(client string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: client,
		// Addr: "localhost:6379",
	})
}