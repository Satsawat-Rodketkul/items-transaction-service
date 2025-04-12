package models

import "time"

type Transaction struct {
	TransactionId string
	UserId        string
	Title         string
	Price         int
	categoty      string
	createDate    time.Time
}

func (Transaction) TableName() string {
	return "transaction"
}
