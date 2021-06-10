package main

import "strconv"

/*
read array of values
check port
configure mysql
return mulitple values from function
*/

func dbConfiguration(properties []string) (string, error) {

	mysqldata := make([]string, 3)
	//check reserved port or not
	port, err := strconv.Atoi(properties[2])
	if port <= 1027 {
		for _, value := range properties {
			mysqldata = append(mysqldata, value)
		}
	}
	return "mysql configured", err
}
