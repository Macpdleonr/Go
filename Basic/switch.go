package main

import {
	"fmt"
}

func main () {
	var fruit string

	fmt.Print("Enter a fruit name: ")
	fmt.Scan(&fruit)

	switch fruit {
	case "apple":
		fmt.Println("You selected an apple.")
	case "banana":
		fmt.Println("You selected a banana.")
	case "orange":
		fmt.Println("You selected an orange.")
	default:
		fmt.Println("Unknown fruit selected.")
		
	}
}