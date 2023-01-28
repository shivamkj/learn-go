package main

import (
	"fmt"
)

func main() {
	// IF STATEMENT
	// if statement executes a block of code if the condition is true
	cartValue, items := 120, 4
	if cartValue > 100 && items > 3 { //parentheses are optional in Go and usually not used
		fmt.Println("You don't need to pay for shipping!")
	}

	// IF, ELSE-IF & ELSE STATEMENT
	// if statement can be followed by an optional else if statement and an optional else statement
	ramScore, rameshScore := 80, 120
	if ramScore == rameshScore {
		fmt.Println("It's a tie!")
	} else if ramScore > rameshScore {
		fmt.Println("Ram wins!")
	} else {
		fmt.Println("Ramesh wins!")
	}

	// SHORT HAND IF STATEMENT
	// if statement can start with a short statement to execute before the condition
	// variables declared by the statement are only in scope until the end of the if
	if ramScore := 120; ramScore > 100 {
		fmt.Println("Ram has scored more than 100!")
	}
	// fmt.Println(ramScore) // this will throw an error as ramScore is not in scope

	// SWITCH STATEMENT
	// switch statement is a shorter way to write a sequence of if-else statements
	// switch statement can be used to compare types, values, expressions
	// In go, behind the scenes, switch statement is converted to a long if-else statement
	var grade string
	switch ramScore {
	case 80, 90: // multiple expressions can be used in a case
		grade = "A+"
		comment := "Excellent!" // Each case block has its own scope
		fmt.Println(comment)
	case 60: // each case condition should be unique
		grade = "A"
		break // break is optional in Go
	case 50: // type of switch condition is int, so case should be same type as well (int)
		grade = "B"
	default: // default case is optional & not necessary to be at the end, only one default case is allowed
		grade = "C"
	}
	// fmt.Println("Comment", comment) // this will throw an error as comment is not in scope
	fmt.Println("Ram's grade is", grade)

	switch { // default switch condition is true
	case ramScore > 90: // boolean expression can be used in a case
	case ramScore > 80: // multiple cases can be provided for the same block
		fmt.Println("Ram scored A+")
	case ramScore > 60:
		fmt.Println("Ram scored A grade")
	default:
		fmt.Println("Ram scored B grade")
	}
}
