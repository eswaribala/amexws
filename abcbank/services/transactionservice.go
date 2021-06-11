package services

import (
	"github.com/amexws/abcbank/models"
)

func CreateTransaction() models.Transaction {
	transaction := models.Transaction{
		25483253,
		"10.30",
		"Parameswari",
		"Bala",
	}
	return transaction
}
