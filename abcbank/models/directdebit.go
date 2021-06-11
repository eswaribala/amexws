package models

type Date struct {
	day   int
	month int
	year  int
}
type DirectDebit struct {
	Payment_Date *Date
}
