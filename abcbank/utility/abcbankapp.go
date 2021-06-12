package main

import (
	"github.com/amexws/abcbank/models"
	"github.com/amexws/abcbank/services"
)

func main() {
	//either withdraw or deposit
	transaction := services.CreateTransaction()
	//implements
	var iTransaction models.ITransaction = &transaction
	iTransaction.DepositMoney(50000)
	directDebit := services.CreateDirectDebit()
	//implements
	iTransaction = &directDebit
	iTransaction.DepositMoney(1000000)

}
