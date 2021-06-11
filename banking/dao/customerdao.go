package dao

import (
	"github.com/amexws/banking/models"
	"math/rand"
	"strconv"
)

func findAllCustomers() models.Customer {

	//create customer instances
	var customers models.Customer

	//for i:=0;i<len(customers);i++ {
	customers = models.Customer{strconv.Itoa(rand.Int()),
		"customer" + (strconv.Itoa(i)), models.Address{"428998", "x", "Chennai", "TN"},
		9952056789, "sample@gmail.com", "test@123"}
	//}

	return customers

}
