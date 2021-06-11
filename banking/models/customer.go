package models

type Address struct {
	door_no     string
	street_name string
	city        string
	state       string
}

type Customer struct {
	account_number int64
	name           string
	address        Address
	contact_number int64
	email          string
	password       string
}
