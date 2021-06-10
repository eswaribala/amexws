package main

import (
	"fmt"
	"reflect"
	"strconv"
)

//custom type

type customers []string

func createBankInfo() {
	//explicitly typed
	//var bankName string="ABC Bank";
	//auto inferred
	bankName := "ABC Bank" //initialization
	bankName = "AMEX"      //reuse
	var city, state = "Chennai", "TN"
	//formatters
	fmt.Printf("The Bank %s in the city %s and state %s\n", bankName, city, state)
	balance := "50000"
	var convBalance, err = strconv.Atoi(balance)
	//data type checking
	fmt.Println(reflect.TypeOf(convBalance).Kind())
	//conditional statements
	if err == nil {
		fmt.Printf("The numerical value of balance=%d\n", convBalance)
	} else {
		fmt.Printf("The associated error%s\n", err)
	}
}

//array of customers
func createCustomers() {
	//custom type usage
	customerData := customers{"Ajay", "Amit", "KK", "Sandeep"}
	//looping through values

	for index, customer := range customerData {
		fmt.Printf("Position is %d and Name=%s ", index, customer)
	}

	//return customers;
}
