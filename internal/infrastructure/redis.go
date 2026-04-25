package infrastructure

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(client string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: client,
	})

	ctx := context.Background()

	// ping check
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("Redis connection failed: %v", err)
	}

	log.Println("Redis connected successfully")
	return rdb, nil
}
