package services

import (
	"github.com/amexws/abcbank/models"
	"math/rand"
)

func CreateTransaction() models.Transaction {
	transaction := models.Transaction{
		rand.Int63n(100000),
		39596,
		"10.30",
		"Parameswari",
		"Bala",
	}
	return transaction
}
