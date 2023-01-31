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
	if ramScoreInner := 120; ramScoreInner > 100 {
		fmt.Println("Ram has scored more than 100!")
	}
	// fmt.Println(ramScoreInner) // this will throw an error as ramScoreInner is not in scope

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

	// SHORT SWITCH STATEMENT
	// switch statement can start with a short statement to execute before the condition
	// variables declared by the statement are only in scope until the end of the switch

	// FALLTHROUGH
	// fallthrough keyword is used to execute the next case block without checking the next case condition
	// used to execute multiple case blocks without checking the condition, should be used at the end of the case block

	switch score, category := 122, "general"; { // default switch condition is true
	case score > 120 && category == "general": // boolean expression can be used in a case
		fallthrough
	case score > 80 && category == "pwd":
		fmt.Println("You are qualified for the next round!")
	case score > 100 && category == "general":
		fallthrough
	case score > 60 && category == "pwd":
		fmt.Println("You need to appear for improvement exam!")
	default:
		fmt.Println("You are not disqualified!")
	}

}
