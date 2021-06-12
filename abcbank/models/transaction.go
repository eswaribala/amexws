package models

type Transaction struct {
	TransactionId int64
	Amount        int64
	Time_Stamp    string
	Sender        string
	Receiver      string
}
