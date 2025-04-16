package main

import (
	"items-transaction-service/internal/config"
	"items-transaction-service/internal/database"
	"items-transaction-service/internal/kafka"
	"items-transaction-service/internal/kafka/consumer"
)

func main() {
	config.LoadEnv()
	database.DBconnection()
	kafka.KafkaConnection()
	consumer.ConsumeMessageTransaction()
}
