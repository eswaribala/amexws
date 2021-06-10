package main

import (
	"fmt"
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

}
