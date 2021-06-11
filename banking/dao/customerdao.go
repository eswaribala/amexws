package dao

import (
	"fmt"
	"github.com/amexws/banking/models"
	"math/rand"
	"strconv"
)

func CreateCustomers() {

	customerList := map[int64]struct {
		name  string
		email string
	}{
		rand.Int63n(100000): {
			"Parameswari",
			"param@gmail.com",
		},
		rand.Int63n(100000): {
			"Bala",
			"bala@gmail.com",
		},
		rand.Int63n(100000): {
			"Vignesh",
			"viki@gmail.com",
		},
	}

	for key, value := range customerList {
		fmt.Printf("\nThe Customer Id is %d and value is %+v", key, value)
	}

}

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
