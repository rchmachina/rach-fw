package redis

import (
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	redisAuth "github.com/rchmachina/rach-fw/internal/repository/redis/auth"
	redisGeneral "github.com/rchmachina/rach-fw/internal/repository/redis/general"
	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	General redisGeneral.RedisGeneralRepository
	Auth    redisAuth.RedisAuthRepository
}

func NewRedisConn(
	client *redis.Client,
	logger logger.Logger,
) *RedisRepository {
	return &RedisRepository{
		General: *redisGeneral.NewRedisGeneral(client, logger),
		Auth:    *redisAuth.NewRedisAuthRepository(client,logger),
	}
}
