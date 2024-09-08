package main

import (
	"github.com/spf13/viper"
	"main/config"
)

func main() {
	config.LoadConfig()
	logger := config.NewLogger(config.DefaultLoggerConfig)
	logger.Info("This is an info message for DEBUG value : ", viper.GetBool("DEBUG"))
	logger.Warning("This is a warning message for services value: ", viper.GetString("SERVICES"))
	logger.Error("This is an error message for ENVIRONMENT value: ", viper.GetString("ENVIRONMENT"))
}
