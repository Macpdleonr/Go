package main

import (
	"fmt"
	"string"
)

func main() {

	var sum int = 0;

	for i := 0; i <= 100; i++ {

		if i%2 != 0 {
			sum = suma + i
		}
	}

	fmt.Println("The sum is:", sum)

	/*for each*/

	myMap := map[string]string{
		"Peru": "Lima",
		"USA":  "Washington DC",
		"UK":   "London",
	}

	for key, value := range myMap {
		fmt.Println("The capital of", key, "is", value)
	}

	/* while */

	var fruit string = "";

	for {

		fmt.Println("Enter a fruit (type 'orange' to exit):")
		fmt.Scan(&fruit)
		fruit = strings.ToLower(fruit)

		if fruit == "orange" {
			fmt.Println("You selected orange, exiting loop.")
			break;
		} else {
			fmt.Println("Current fruit is:", fruit)
		}
	}

}