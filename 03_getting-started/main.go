package main

import (
	"fmt"
	"os"
)

func main() {
	// TYPE CONVERSION syntax: T(v) [type(value)]
	var distance int = 100
	var time float64 = 25.5
	var speed float64 = float64(distance) / time
	fmt.Println(speed)
	score := 100
	var increase = 5.5
	score = int(float64(score) * increase) // avoids loss of precision during type conversion
	fmt.Println(score)

	// COMMAND LINE ARGUMENTS
	fmt.Println(os.Args) // get the command line arguments passed to the program,
	// first argument is the program path, the rest are the arguments passed to the program
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}

	// Printf vs Println
	// Printf formats and prints the string and replaces the placeholders with the values passed
	// Println prints the string with a new line at the end
	fmt.Printf("Hello, %s! You scored %d points.\nCongratulation %[1]v! \n", "Shivam", 100)

	// 	CONSTANTS
	// 	constants are declared using the const keyword
	// 	constants can be character, string, boolean or numeric values
	// 	cannot be declared using the := syntax
	const pi = 3.14
	const e = 2.71828
	// pi = 3.14159 // error: cannot assign to pi

	// IOTA
	// iota is a predeclared identifier representing the untyped integer
	// ordinal number of the current const specification in a (usually parenthesized) const declaration.
	// It is zero-indexed.
	const (
		_  = iota // 0
		KB        // 1
		MB        // 2
		GB        // 3
		TB        // 4
		PB        // 5
	)
	fmt.Println(KB, MB, GB, TB, PB)

}
