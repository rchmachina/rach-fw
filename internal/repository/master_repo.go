package repository

import (
	"github.com/rchmachina/rach-fw/internal/infrastructure/logger"
	db "github.com/rchmachina/rach-fw/internal/repository/db"
	noSql "github.com/rchmachina/rach-fw/internal/repository/redis"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Repository struct {
	Sql   db.Repositories
	NoSql noSql.RedisRepository
}

func NewMasterRepo(
	conns *gorm.DB,
	client *redis.Client,
	logger logger.Logger,
) *Repository {
	return &Repository{
		Sql:   db.NewDbRepositoryMaster(conns, logger),
		NoSql: *noSql.NewRedisConn(client, logger),
	}
}
