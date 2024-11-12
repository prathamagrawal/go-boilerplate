package main

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"main/config"
	"main/utils"
)

func PublishUtil(logger *logrus.Logger, data map[string]interface{}, exchange string, routingKey string) {
	var ctx = context.Background()
	channelObject := utils.GetQueueConnection()
	body, err := json.Marshal(data)
	if err != nil {
		utils.FailOnError(err, "Failed to convert map to JSON")
	}
	err = channelObject.PublishWithContext(ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	utils.FailOnError(err, "Failed to publish a message")
	logger.Info(" [x] Sent JSON message: %s\n", data)
}

func main() {
	config.LoadConfig()
	logger := config.NewLogger(config.DefaultLoggerConfig)
	data := map[string]interface{}{
		"message": "Hello World!",
		"user":    "john_doe",
		"id":      1234,
	}
	PublishUtil(logger, data, "test_exchange1", "routing.key1.1")

}
