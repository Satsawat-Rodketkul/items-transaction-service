package consumer

import (
	"context"
	"encoding/json"
	"items-transaction-service/internal/config"
	"items-transaction-service/internal/database/models"
	"items-transaction-service/internal/database/repositories"
	"items-transaction-service/internal/kafka"
	"log"
)

func ConsumeMessageTransaction(msg []byte) {
	topic := config.GetValue("KAFKA_TOPIC")

	err := kafka.Consumer.Consume(context.Background(), []string{topic}, NewConsunerGroupHandler())
	if err != nil {
		log.Fatal(err)
	}

	defer kafka.Consumer.Close()

	transaction := models.Transaction{}

	err = json.Unmarshal(msg, &transaction)
	if err != nil {
		log.Println("Unmarshal error:", err)
	}

	log.Printf("Consumed data: %+v", transaction)

	err = repositories.TransactionBuyItem(&transaction)
	if err == nil {
		log.Println("Insert transaction buy item success")
	}
}
