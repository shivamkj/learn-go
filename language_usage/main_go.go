package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// COMMAND LINE ARGUMENTS
	fmt.Println(os.Args) // get the command line arguments passed to the program,
	// first argument is the program path, the rest are the arguments passed to the program
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}

	// SHADOWING
	// shadowing is the process of declaring a new variable with the same name as a previous variable
	// both variables have different scopes & are different variables
	var myShadow int = 10
	if myShadow := 100; myShadow > 50 {
		fmt.Println("Shadowing", myShadow)
	}
	fmt.Println("Shadowing", myShadow)

	// Printf vs Println
	// Printf formats, replaces the placeholders with the values passed and prints the string
	// Println prints the string with a new line at the end
	fmt.Printf("Hello, %s! You scored %d (%[2]T) points.\nCongratulation %[1]v! \n", "Shivam", 100)

	// TIME PACKAGE
	// provides functions to get the current time, parse time, format time, etc.
	fmt.Println("Today is", time.Now().Weekday())
}
