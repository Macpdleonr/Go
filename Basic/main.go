package main

import "fmt"

var functions = map[string] func(int, int) int{

	"sum": func(a int, b int) int { return a + b },
	
	"rest": func(a int, b int) int { return a - b },

	"multiply": func(a int, b int) int { return a * b },

	"divide": func(a int, b int) int { return a / b },
}

func presentResult(operation string, a int, b int){

	f, exists := functions[operation]

	if !exists {
		fmt.Println("Operation ", operation, " not found")
		return
	}
	
	fmt.Println("For a = ", a, " and b = ", b, " the ", operation," result is : ", f(a,b))
}

func main() {

	presentResult("sum",3, 4)
	presentResult("sum",10, 20)
	presentResult("sum",-55, 83)

	presentResult("rest",5, 2)
	presentResult("rest",100, 50)
	presentResult("rest",80, 25)

	presentResult("multiply",3, 4)
	presentResult("multiply",10, 20)
	presentResult("multiply",-5, 6)

	presentResult("divide",8, 2)
	presentResult("divide",100, 5)
	presentResult("divide",81, 9)

}