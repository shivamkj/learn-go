package main

import (
	"fmt"
	"path"
)

var toplevel = "Hello, playground" // declare a variable at the package level
// only variables assignment can be used, short declaration is not allowed at the package level

func main() {
	// 3 ways to declare a variable
	// 		1. var name type = expression 	// type is explicitly specified
	// 		2. var name = expression 	// type is inferred from the expression, can't be used without initialization
	// 		3. name := expression	// short declaration
	fmt.Println(toplevel)

	var speed int       // DECLARATION: declare a variable
	fmt.Println(speed)  // print 0 (zero value, as not initialized)
	speed = 100         // INITIALISATION: assign a value to the variable
	var distance = 1000 // DECLARATION & INITIALISATION: declare and initialize a variable
	fmt.Println(speed, distance)

	var force, power float64 // declare multiple variables
	force, power = 1.5, 2.5  // assign multiple values to multiple variables
	fmt.Println(force, power)

	var direaction string = "left" // declare and initialize a variable
	fmt.Println(direaction)

	one, two := 1, 2 // declare with short declaration and initialize two variables
	fmt.Println(one, two)

	_, path := path.Split("css/main.css") // use the blank identifier to ignore the first return value
	fmt.Println(path)

	// use variable assignment to declare and/or initialize multiple variable closely related
	var (
		न      string = "शिवम" // unicode characters are allowed in variable names & string literals
		age    int    = 32
		height int    = 180
		// declare variables without initializing it
		weight  int
		address string
	)
	fmt.Println(न, age, height, weight, address)

	// SWAPPING two variables
	var condition1, condition2 bool = true, false
	condition1, condition2 = condition2, condition1

	// VARIABLE DECLARATION RULES
	// 1. A variable must be declared before it is used with a unique name
	// 2. Must start with a letter or an underscore & can't start with a number
	// 3. A variable name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
	//	- no special characters & space allowed
	// 4. Variable names are case-sensitive (age, Age and AGE are three different variables)
	// 5. A variable cannot be declared more than once in the same scope
	// 6. A variable cannot be declared with a reserved keyword, functions, identifiers (package, func, var, int, string, bool, etc.)
	//	- , float64, bool, string, complex64, complex128, byte, rune, error, uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, complex64, complex128, bool, byte, rune, string, error, uintptr
	// Some Examples of invalid Names: 3make, !happy, len, func, package, int

	// var testing int = 100 // ERROR: go does not allow unused variables

}
