package dao

import (
	"github.com/amexws/banking/models"
	"math/rand"
	"strconv"
)

func FindAllCustomers() models.Customer {

	//create customer instances
	var customers models.Customer

	//for i:=0;i<len(customers);i++ {
	customers = models.Customer{rand.Int63n(1000000),
		"customer" + (strconv.Itoa(rand.Int())), models.Address{"428998", "x", "Chennai", "TN"},
		9952056789, "sample@gmail.com", "test@123"}
	//}

	return customers

}
