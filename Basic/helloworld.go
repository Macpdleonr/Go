package main
import "fmt"

func main() {
	
	fmt.Println("Hello, World!")

	var name string  = "Mac"

	fmt.Printf("Hello, %s!\n", name)

	var lastname = "Testing"

	fmt.Printf("Hello, %s %s!\n", name, lastName)

	petname := "Max"

	fmt.Printf("Hello, %s!\n", petname)

	/* Numeric */
	var yearnow int16 = 2024
	fmt.Printf("The year is %d\n", yearnow)
	var lessyear int8 = 123
	fmt.Printf("The year is %d",lessyear)
	oldyear := 33
	fmt.Printf("Old year is %d\n", lessyear)

	/* Array */

	var frutsList = [4] string {"Apple", "Banana", "Grapes", "Orange"}

	fmt.Println(frutsList[1])

	// countryList := [3] string {"USA", "UK", "India"}
	countryList := [] string {"USA", "UK", "India"}
	fmt.Println(countryList)

	countryList[0] = "Peru"

	fmt.Println(countryList)

	append(countryList, "Spain")

	fmt.Println(countryList)

	countryList = append(countryList, "Portugal")

	fmt.Println(countryList)

	// Rangs

	countryList2 := countryList[1:3]

	fmt.Println(countryList2)

	countryList3 := countryList[2:]

	fmt.Println(countryList3)

	countryList4 := countryList[:2]

	fmt.Println(countryList4)


	}

