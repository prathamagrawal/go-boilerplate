package utils

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	"main/config"
)

func GetQueueConnection() *amqp.Channel {
	logger := config.NewLogger(config.DefaultLoggerConfig)
	amqpConnectionString := viper.GetString("AMQP_CONNECTION")
	logger.Info("Getting Queue Connection on:", amqpConnectionString)
	conn, err := amqp.Dial(amqpConnectionString)
	if err != nil {
		FailOnError(err, "Failed to connect to %s")
	}
	ch, err := conn.Channel()
	if err != nil {
		FailOnError(err, "Failed to open a channel")
		defer ch.Close()
	}
	return ch
}
