package models

import (
	"fmt"
)

type ITransaction interface {
	DepositMoney(Amount int64)
	WithdrawMoney(Amount int32)
}

//methods
func (transaction *Transaction) DepositMoney(Amount int64) {
	fmt.Printf("\nTransaction %+v amount deposited %d", transaction, Amount)
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
