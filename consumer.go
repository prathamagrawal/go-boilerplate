package main

import (
	"github.com/spf13/viper"
	"log"
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

func main() {

	EXCHANGES := []map[string]interface{}{
		{"name": "test_exchange1", "type": "direct", "durable": true},
		{"name": "test_exchange2", "type": "direct", "durable": true},
	}

	queues := []Queue{
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

	config.LoadConfig()
	logger := config.NewLogger(config.DefaultLoggerConfig)
	logger.Info("Starting consumer with the following pubsub broker: ", viper.GetString("PUBSUB"))
	channelObject := utils.GetQueueConnection()

	// Declaring Exchanges with the configuration provided above:
	for _, exchange := range EXCHANGES {
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
		queues[0].Name,
		false,
		false,
		true,
		false,
		nil,
	)
	if err != nil {
		utils.FailOnError(err, "Failed to create queue")
	}
	for _, queue := range queues {
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
			log.Printf(" [x] %s", d.Body)
		}
	}()

	logger.Info(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever

}