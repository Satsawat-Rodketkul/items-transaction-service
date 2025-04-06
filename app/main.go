package main

import (
	"items-transaction-service/internal/config"
	"items-transaction-service/internal/database"
)

func main() {
	config.LoadEnv()
	database.DBconnection()
}
