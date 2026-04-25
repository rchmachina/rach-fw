package redis

import (
	"context"
	"time"

	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"github.com/redis/go-redis/v9"
)

type RedisGeneralRepository struct {
	client *redis.Client
	logger logger.Logger
}

func NewRedisGeneral(
	client *redis.Client,
	logger logger.Logger,
) *RedisGeneralRepository {
	return &RedisGeneralRepository{
		client: client,
		logger: logger,
	}
}

func (r *RedisGeneralRepository) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) error {


	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {

		r.logger.Error("failed to set redis key",
			logger.Field{Key: "key", Value: key},
			logger.Field{Key: "error", Value: err},
		)
		return err
	}

	return nil
}

func (r *RedisGeneralRepository) Get(
	ctx context.Context,
	key string,
) (string, error) {

	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil // key not found
		}

		r.logger.Error("failed to get redis key",
			logger.Field{Key: "key", Value: key},
			logger.Field{Key: "error", Value: err},
		)
		return "", err
	}

	return val, nil
}

func (r *RedisGeneralRepository) Delete(
	ctx context.Context,
	key string,
) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		r.logger.Error("failed to delete redis key",
			logger.Field{Key: "key", Value: key},
			logger.Field{Key: "error", Value: err},
		)
		return err
	}

	return nil
}

func (r *RedisGeneralRepository) TTL(
	ctx context.Context,
	key string,
) (time.Duration, error) {

	return r.client.TTL(ctx, key).Result()
}


