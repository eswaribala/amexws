package main

import (
	"github.com/amexws/abcbank/models"
	"github.com/amexws/abcbank/services"
)

func main() {
	//either withdraw or deposit
	transaction := services.CreateTransaction()
	var iTransaction models.ITransaction = &transaction
	iTransaction.DepositMoney(50000)
	directDebit := services.CreateDirectDebit()
	iTransaction = &directDebit
	iTransaction.DepositMoney(1000000)

}
