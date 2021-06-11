package models

import "fmt"

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

type PlatinumCustomer struct {
	Customer
	Limit int64
	Offer int32
}

/**

Method has receiver type
*/
//update the address
//call by reference
func (customer *Customer) UpdateAddress(address *Address) {

	(*customer).AddressRef = *address
}

//update the contact no
//call by value
//receiver type,parameter list,return type
func (customer Customer) UpdateContactNo(contact_no int64) Customer {

	customer.Contact_Number = contact_no
	fmt.Printf("\nInternal Customer Details %+v", customer)
	return customer
}

//method overriding

func (customer *Customer) Withdraw() {
	fmt.Printf("\nCustomer Data %+v", customer)
}

func (customer *PlatinumCustomer) Withdraw() {
	fmt.Printf("\nCustomer Data %+v", customer)
}
