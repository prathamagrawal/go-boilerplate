package main

import (
	"context"
	"main/config"
	"main/utils"
)

func main() {
	logger := config.NewLogger(config.DefaultLoggerConfig)
	ctx := context.Background()
	config.LoadConfig()
	rdb := utils.GetRedisConnection()
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		logger.Error("redis get key failed")
	}
	logger.Info("Key found with value: ", val)
}
