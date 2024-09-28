package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"main/config"
	"main/utils"
)

type Binding struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
}
type Queue struct {
	Name     string    `json:"name"`
	Durable  bool      `json:"durable"`
	Bindings []Binding `json:"bindings"`
}

var Exchanges = []map[string]interface{}{
	{"name": "test_exchange1", "type": "direct", "durable": true},
	{"name": "test_exchange2", "type": "direct", "durable": true},
}

var Queues = []Queue{
	{
		Name:    "go_queue",
		Durable: true,
		Bindings: []Binding{
			{
				Exchange:   "test_exchange1",
				RoutingKey: "routing.key1.1",
			},
			{
				Exchange:   "test_exchange1",
				RoutingKey: "routing.key1.2",
			},
			{
				Exchange:   "test_exchange2",
				RoutingKey: "routing.key2.1",
			},
		},
	},
}

func main() {

	config.LoadConfig()
	logger := config.NewLogger(config.DefaultLoggerConfig)
	logger.Info("Starting consumer with the following pubsub broker: ", viper.GetString("PUBSUB"))
	channelObject := utils.GetQueueConnection()

	for _, exchange := range Exchanges {
		exchange_name := exchange["name"].(string)
		exchange_type := exchange["type"].(string)
		exchange_durable := exchange["durable"].(bool)
		err := channelObject.ExchangeDeclare(
			exchange_name,
			exchange_type,
			exchange_durable,
			false, // auto-deleted
			false, // internal
			false, // no-wait
			nil,   //any extra arguments
		)
		if err != nil {
			utils.FailOnError(err, "Failed to create exchange")
		}
	}

	queueObject, err := channelObject.QueueDeclare(
		Queues[0].Name,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		utils.FailOnError(err, "Failed to create queue")
	}
	for _, queue := range Queues {
		for _, binding := range queue.Bindings {
			err := channelObject.QueueBind(queueObject.Name, binding.RoutingKey, binding.Exchange, false, nil)
			if err != nil {
				utils.FailOnError(err, "Failed to bind to exchange"+binding.Exchange)
			}
		}
	}

	msgs, err := channelObject.Consume(
		queueObject.Name,
		"go_consumer",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			messageData := make(map[string]interface{})
			if err := json.Unmarshal(d.Body, &messageData); err != nil {
				logger.Errorf("Failed to decode JSON: %s", err)
				continue
			}

			logger.Infof("Item: %+v received on routing key %s with exchange %s", messageData, d.RoutingKey, d.Exchange)
			if d.RoutingKey == "routing.key1.1" {
				// do something
				logger.Info("Message received on routing key routing.key1.1")
			} else if d.RoutingKey == "routing.key1.2" {
				logger.Error("Message received on routing key routing.key1.2")
			} else {
				// do something
				logger.Error("Message received on routing key routing.key2.1")
			}

		}
	}()

	logger.Info(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}
