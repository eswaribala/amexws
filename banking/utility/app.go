package main

import (
	//"fmt"
	dao "github.com/amexws/banking/dao"
	//models "github.com/amexws/banking/models"
	//"math/rand"
)

func main() {

	//dao.CreateCustomers()
	dao.CreateCustomerList()

	/*
		fmt.Println(dao.FindAllCustomers())
		fmt.Printf("\nCustomer Data %+v", dao.FindAllCustomers())
		//partial customer struct

		customerObj := dao.FindAllCustomers()
		//inline structure
		partial_customer := struct {
			account_no int64
			name       string
		}{
			account_no: rand.Int63(),
			name:       customerObj.Name,
		}

		fmt.Printf("\nPartial Customer Data%+v", partial_customer)

		//update the address
		address := models.Address{"487658", "Gandhi st", "Bangalore", "KN"}
		fmt.Printf("\nAddress %+v", address)
		// how do we invoke method?
		//use always instance to access method which is receiver type
		customer_response := dao.FindAllCustomers()
		customer_response.UpdateAddress(&address)

		fmt.Printf("\nAfter Address Update Customer Data %+v", customer_response)

		customer_response = customer_response.UpdateContactNo(555555500)
		fmt.Printf("\nAfter Contact Number Update Customer Data %+v", customer_response)

		//methos overriding
		customer_response.Withdraw();
		platinum_customer:=models.PlatinumCustomer{
			customer_response,
			500000,
			10,
		}
		platinum_customer.Withdraw();
	*/
}
