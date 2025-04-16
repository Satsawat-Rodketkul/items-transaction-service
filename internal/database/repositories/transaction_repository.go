package repositories

import (
	"items-transaction-service/internal/database"
	"items-transaction-service/internal/database/models"
	"log"
)

func TransactionBuyItem(transaction *models.Transaction) error {
	tx := database.DB.Create(transaction)
	if tx.Error != nil {
		log.Println("Error insert transaction buy item:", tx.Error)
	}
	return tx.Error
}
