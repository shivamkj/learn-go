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
	score = int(float64(score) * increase)
	fmt.Println(score)

	// COMMAND LINE ARGUMENTS
	fmt.Println(os.Args) // get the command line arguments passed to the program,
	// first argument is the program path, the rest are the arguments passed to the program
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
}
