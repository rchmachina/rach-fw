package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	model "github.com/rchmachina/rach-fw/internal/dto/model/jwt"
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	"github.com/rchmachina/rach-fw/internal/utils/helper"
	"github.com/redis/go-redis/v9"
)

type RedisAuthRepository struct {
	client *redis.Client
	logger logger.Logger
}

func NewRedisAuthRepository(
	client *redis.Client,
	logger logger.Logger,
) *RedisAuthRepository {
	return &RedisAuthRepository{
		client: client,
		logger: logger,
	}
}

func (r *RedisAuthRepository) IncrementLoginAttempt(
	ctx context.Context,
	email string,
	expiration time.Duration,
) (int64, error) {

	key := helper.BuildKey("Auth:LoginAttemp:", email)
	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if count == 1 {
		r.client.Expire(ctx, key, expiration)
	}

	return count, nil
}

func (r *RedisAuthRepository) UpdateLogin(
	ctx context.Context,
	model model.TokenValue,
	value string,
	expiration time.Duration,
) error {

	searchKey := helper.BuildKey("Auth:login:", model.Email)
	count, err := r.CountLoginByPrefix(ctx, searchKey)
	if err != nil {
		return fmt.Errorf("there is something wrong")
	}
	if count > 3 {
		return fmt.Errorf("please log out you account in another device first")
	}

	val, err := json.Marshal(model)
	if err != nil {
		return fmt.Errorf("there is something wrong")
		// return err
	}

	keyToken := helper.BuildKey(searchKey, value)
	err = r.client.Set(ctx, keyToken, val, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisAuthRepository) RevokeRefreshToken(
	ctx context.Context,
	key string,
	value string,
) error {

	searchKey := helper.BuildKey("Auth:login:", key)
	keyToken := helper.BuildKey(searchKey, value)
	return r.client.Del(ctx, keyToken).Err()

}

func (r *RedisAuthRepository) IsTokenExists(
	ctx context.Context,
	key string,
) (*model.TokenValue, error) {

	pattern := "*" + key + "*"

	iter := r.client.Scan(ctx, 0, pattern, 0).Iterator()
	var result model.TokenValue
	for iter.Next(ctx) {
		key := iter.Val()

		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}
		if val != "" {
			err = json.Unmarshal([]byte(val), &result)
			if err != nil {
				return nil, err
			}
			break
		}

	}
	return &result, nil
}

func (r *RedisAuthRepository) CountLoginByPrefix(ctx context.Context, prefix string) (int, error) {
	var (
		cursor uint64
		count  int
	)

	pattern := prefix + "*"

	for {
		keys, nextCursor, err := r.client.Scan(ctx, cursor, pattern, 100).Result()
		if err != nil {
			return 0, err
		}

		count += len(keys)
		cursor = nextCursor

		if cursor == 0 {
			break
		}
	}

	return count, nil
}

func (r *RedisAuthRepository) RemoveLoginAttempt(
	ctx context.Context,
	email string,
	expiration time.Duration,
) error {
	key := helper.BuildKey("Auth:LoginAttemp:", email)
	return r.client.Del(ctx, key).Err()
}

func (r *RedisAuthRepository) GetLoginAttempt(
	ctx context.Context,
	email string,
) (int64, error) {
	key := helper.BuildKey("Auth:LoginAttemp:", email)
	val, err := r.client.Get(ctx, key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	return val, nil
}
