package main

import (
	"fmt"
	"items-transaction-service/internal/config"
	"items-transaction-service/internal/kafka"

	"github.com/IBM/sarama"
)

func main() {
	config.LoadEnv()
	//database.DBconnection()
	kafka.KafkaConnection()
	defer kafka.Consumer.Close()
	defer kafka.Producer.Close()

	topic := config.GetValue("KAFKA_TOPIC")

	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Hello World kafka producer"),
	}
	partition, offset, err := kafka.Producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("partition: %v, offset: %v", partition, offset)

	partitionConsumer, err := kafka.Consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	fmt.Println("Consumer start.")
	for {
		select {
		case err := <-partitionConsumer.Errors():
			fmt.Println(err)
		case msg := <-partitionConsumer.Messages():
			fmt.Println(string(msg.Value))
		}
	}
}
