package main

import (
	"fmt"
	"github.com/amexws/day2/utility"
	"os"
)

//main frame is created stack
func main() {

	createBankInfo()
	for index, customer := range createCustomers() {
		fmt.Printf("Position is %d and Name=%s ", index, customer)
	}
	//readTransactionData();
	generateIds()
	result, err := dbConfiguration(os.Args)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = getDocuments(1246546, "Passport")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)

	}
	result, err = getDocuments(1246546, "Passport", "Driving license")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)

	}

	result, err = getDocuments(1246546)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)

	}

	//invoke function from day2 utility
	customerList := customers{"Param", "Bala", "Vignesh", "Sanjay"}
	//byte slicing
	fmt.Println([]byte(utility.ToString(customerList)))

}
