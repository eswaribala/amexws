package main

import (
	"fmt"
	"math/rand"
)

func generateIds() {
	//create slice
	ids := make([]int32, 50)
	//assign values
	for id := range ids {
		//ids[id]=rand.Int31();
		fmt.Println(id)
		ids = append(ids, rand.Int31())
	}
	//loop over
	for index, id := range ids {
		fmt.Println(index, "=>", id)
	}

}
