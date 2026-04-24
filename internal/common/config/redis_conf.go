package config

import (
	"context"
	"log"
	"strconv"

	redis "github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client
var RedisCtx = context.Background()

const (
	RedisHost     = "redis.host"
	RedisProt     = "redis.port"
	RedisDatabase = "redis.database"
)

type RedisConfig struct {
	Host     string
	Port     int
	Database int
}

func NewRedisInstance() error {
	redisConfig := &RedisConfig{
		Host:     viper.GetString(RedisHost),
		Port:     viper.GetInt(RedisProt),
		Database: viper.GetInt(RedisDatabase),
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + strconv.Itoa(redisConfig.Port),
		DB:       redisConfig.Database,
		Password: "",
	})

	_, err := RedisClient.Ping(RedisCtx).Result()

	if err != nil {
		log.Fatalf("[Redis connect ERROR]: %v \n", err)
	}
	return err
}
