package main 

import "fmt"

/* Maps and Dict*/

func maps() {
	myMap := map[string]string{
		"Peru": "Lima",
		"USA":  "Washington DC",
		"UK":   "London",
	}

	fmt.Println("The map of country is:", myMap)

	fmt.Println("The capital of USA is:", myMap["USA"])

	delete(myMap, "UK")

	fmt.Println("The map of country is:", myMap)
}
