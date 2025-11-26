package main

import "fmt"



func presentResultSum( f func(int, int)int, a int, b int) {
	fmt.Println("For a = ", a," and b = ", b," the sum is : ", sum(a, b))
}

func presentResultRest( f func(int, int)iSnt, a int, b int) {
	fmt.Println("For a = ", a," and b = ", b," the rest is : ", rest(a, b))
}

func presentResult(operation string, a int, b int){
	sum := func(a int, b int) int {
		return a + b
	}

	rest := func(a int, b int) int {
		return a - b
	}

	result := 0

	if operation == "sum" {
		result = sum(a, b)
	} else if operation == "rest" {
		result = rest(a, b)
	}

	fmt.Println("For a = ", a, " and b = ", b, " the ", operation," result is : ", result)
}

func main() {

	



	presentResultSum("sum",3, 4)
	presentResultSum("sum",10, 20)
	presentResultSum("sum",-55, 83)

	presentResultRest("rest",5, 2)
	presentResultRest("rest",100, 50)
	presentResultRest("rest",80, 25)

}