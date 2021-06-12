package models

import (
	"fmt"
	"github.com/amexws/abcbank/storeorm"
	"github.com/amexws/abcbank/stores"
	"log"
)

type ITransaction interface {
	DepositMoney(Amount int64)
	WithdrawMoney(Amount int32)
}

//methods
func (transaction *Transaction) DepositMoney(Amount int64) {
	fmt.Printf("\nTransaction %+v amount deposited %d", transaction, Amount)
	result, err := stores.CreateTransaction(transaction.TransactionId, transaction.Amount, transaction.Time_Stamp, transaction.Sender, transaction.Receiver)
	if err != nil {
		log.Fatal("User Not inserted....")

	}
	fmt.Println("\nRecord Inserted", result)

	storeorm.SaveTransaction(transaction.TransactionId, transaction.Amount, transaction.Time_Stamp, transaction.Sender, transaction.Receiver)
	fmt.Println("\nRecord Inserted", result)
	stores.GetAllTransactions()

}
func (directDebit *DirectDebit) DepositMoney(Amount int64) {
	fmt.Printf("\nDirect Debit %+v amount deposited %d", directDebit, Amount)
}
func (externalTransaction *ExternalTransaction) DepositMoney(Amount int64) {

}

func (transaction *Transaction) WithdrawMoney(Amount int32) {

}
func (directDebit *DirectDebit) WithdrawMoney(Amount int32) {

}
func (externalTransaction *ExternalTransaction) WithdrawMoney(Amount int32) {

}
