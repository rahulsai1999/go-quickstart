package src

import "fmt"

//Variables -> initializing variables
func Variables() {
	var x int // declaration
	x = 3     // assignment
	y := 4    // short declaration - infers the type,declare and assign

	sum, prod := Functions(x, y)
	fmt.Println("Sum:", sum, " Prod:", prod)
	return
}

//Functions -> defining functions
func Functions(x, y int) (sum, prod int) {
	// function parameters are passed like this
	// x, y are the input arguments of type int
	// sum, prod are the output signatures of type int
	return x + y, x * y
}

//FunctionsNamedReturn -> defining functions with named returns
func FunctionsNamedReturn(x, y int) (z int) {
	z = x * y
	return // z is implicit as only single parameter is present
}
