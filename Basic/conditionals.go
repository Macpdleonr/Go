package main
import "fmt"

func conditionals() {
	
	var yearOld int = 32;

	var permission bool = false

	if yearOld < 18 && permission {
		fmt.Println("You cant travel alone")
	} else if yearOld >= 18 && permission{
		fmt.Println("You can travel")
	} else if yearOld >= 18 && !permission {
		fmt.Println("You cant travel without permission")
	}

}
