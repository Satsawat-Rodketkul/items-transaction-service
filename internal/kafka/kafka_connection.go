package kafka

import (
	"fmt"
	"items-transaction-service/internal/config"
	"log"

	"github.com/IBM/sarama"
)

var Consumer sarama.ConsumerGroup

func KafkaConnection() {
	kafkaHost := config.GetValue("KAFKA_HOST")
	kafkaPort := config.GetValue("KAFKA_PORT")
	kafkaGroup := config.GetValue("KAFKA_TXN_GROUP")

	server := []string{fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)}

	consumer, err := sarama.NewConsumerGroup(server, kafkaGroup, nil)
	if err != nil {
		log.Fatal("Failed to connect kafka server:", err)
	}
	Consumer = consumer

	log.Print("Connect to kafka server success")
}
