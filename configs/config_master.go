// Package config is dictionary of all config struct
package configs

import (
	"log"
	"os"
	"strconv"
	"time"
)

var osGetenv = os.Getenv

type Configs struct {
	DbConf        string
	RedisConf     string
	Port          string
	IsProduction  bool
	AccesTokenKey string
	TTlRedis      TTlRedis
}
type TTlRedis struct {
	TTLAttemp       time.Duration
	TTLRefreshToken time.Duration
	TTLAccessToken  time.Duration
}

func LoadConfig() *Configs {
	isProduction := GetConfig("IS_PRODUCTION")
	isProductionBool, err := strconv.ParseBool(isProduction)
	if err != nil {
		isProductionBool = false
	}

	///redis ttl
	TtlAttemp := GetConfig("TTL_ATTEMP")

	TtlAttempTime, err := time.ParseDuration(TtlAttemp)
	if err != nil {
		log.Fatal(err)
	}
	TtlRefreshToken := GetConfig("TTL_REFRESH_TOKEN")
	TtlRefreshTokenTime, err := time.ParseDuration(TtlRefreshToken)
	if err != nil {
		log.Fatal(err)
	}

	TtlAccessToken := GetConfig("TTL_ACCESS_TOKEN")

	TtlAccessTokenTime, err := time.ParseDuration(TtlAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	return &Configs{
		DbConf:        osGetenv("DSN_DB"),
		RedisConf:     osGetenv("REDIS_CLIENT"),
		Port:          osGetenv("PORT"),
		IsProduction:  isProductionBool,
		AccesTokenKey: osGetenv("ACCESS_TOKEN"),
		TTlRedis: TTlRedis{
			TTLAttemp:       TtlAttempTime,
			TTLRefreshToken: TtlRefreshTokenTime,
			TTLAccessToken:  TtlAccessTokenTime,
		},
	}
}

func GetConfig(key string) string {
	return osGetenv(key)
}
