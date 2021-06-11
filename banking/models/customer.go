package models

type Address struct {
	Door_No     string
	Street_Name string
	City        string
	State       string
}

type Customer struct {
	Account_Number int64
	Name           string
	AddressRef     Address
	Contact_Number int64
	Email          string
	Password       string
}
