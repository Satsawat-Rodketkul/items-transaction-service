package models

import "time"

type Transaction struct {
	TransactionId string
	UserId        string
	Title         string
	Price         float64
	Categoty      string
	CreateDate    time.Time
}

func (Transaction) TableName() string {
	return "transaction"
}
