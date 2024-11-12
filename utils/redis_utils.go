package utils

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"main/config"
)

func GetRedisConnection() *redis.Client {
	redisConnectionUrl := viper.GetString("REDIS_URL")
	logger := config.NewLogger(config.DefaultLoggerConfig)
	logger.Info("Connecting to redis on URL: ", redisConnectionUrl)
	conn, err := redis.ParseURL(redisConnectionUrl)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(conn)
	return rdb

}
