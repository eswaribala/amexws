package main

import (
	"fmt"
	"github.com/amexws/banking/dao"
)

func main() {
	fmt.Println(dao.findAllCustomers())
}
