package kafka

import (
	"fmt"
	"items-transaction-service/internal/config"
	"log"

	"github.com/IBM/sarama"
)

var Consumer sarama.Consumer
var Producer sarama.SyncProducer

func KafkaConnection() {
	kafkaHost := config.GetValue("KAFKA_HOST")
	kafkaPort := config.GetValue("KAFKA_PORT")

	server := []string{fmt.Sprintf("%s:%s", kafkaHost, kafkaPort)}

	consumer, err := sarama.NewConsumer(server, nil)
	if err != nil {
		log.Fatal("Failed to connect kafka server:", err)
	}
	Consumer = consumer

	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		log.Fatal("Failed to connect kafka server:", err)
	}
	Producer = producer

	log.Print("Connect to kafka server success")
}
