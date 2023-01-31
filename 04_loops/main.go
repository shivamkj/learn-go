package main

import "fmt"

func main() {
	// FOR LOOP
	// for init; condition; post {}
	// init: executed before the first iteration (optional)
	// condition: evaluated before every iteration (optional - if not provided, default is always true, infinite loop)
	// post: executed at the end of every iteration (optional)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// Write the same loop without init & post
	index := 0
	for index < 3 {
		fmt.Println("Think Different !", index)
		index++
	}

	// INIFINITE LOOP
	// for {
	// 	fmt.Println("Will run forever!")
	// }

	// BREAK & CONTINUE
	// break: used to exit the loop
	// continue: used to skip the current iteration
	for i := 0; i < 4; i++ {
		if i == 1 {
			continue
		}
		for inner := 0; inner < 3; inner++ { // Nested loop
			if inner == 2 {
				break
			}
			fmt.Println("CONTINUE", inner, i)
		}
		fmt.Println("BREAK", i)
	}

	// LOOPING OVER COLLECTIONS
	var topCities = []string{"Mumbai", "Delhi", "Bangalore", "Hyderabad", "Chennai"}
	for i := 0; i < len(topCities); i++ { // using for loop
		fmt.Println(i, topCities[i])
	}
	// SYNTAX for index (current index element), value (current index value, optional) := range (keyword) collection {}
	for index, value := range topCities { // using range (recommended, index can be ignored by using _)
		fmt.Println(index, value) //
	}

}
