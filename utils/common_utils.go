package utils

import (
	"main/config"
)

func FailOnError(err error, msg string) {
	logger := config.NewLogger(config.DefaultLoggerConfig)
	if err != nil {
		logger.Panicf("%s: %s", msg, err)
	}
}
