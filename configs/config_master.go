// Package config is dictionary of all config struct
package configs

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var osGetenv = os.Getenv

type Configs struct {
	Db    *gorm.DB
	Redis *redis.Client
	Port  string
}

func LoadConfig() (*Configs, error) {
	getDb, err := NewDB(osGetenv("DSN_DB"))
	if err != nil {
		log.Panic(err)
	}
	getRedis := NewRedisClient(osGetenv("REDIS_CLIENT"))

	return &Configs{
		Db:    getDb,
		Redis: getRedis,
		Port:  GetConfig("PORT"),
	}, nil
}

func GetConfig(key string)string{
	return osGetenv(key)
}
