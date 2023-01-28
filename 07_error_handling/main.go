package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		toConvert := os.Args[1]
		num, err := strconv.Atoi(toConvert)
		if err != nil {
			fmt.Printf("Passed value is not an integer value %v", err)
			os.Exit(1)
		}
		fmt.Printf("The number is %d\n", num)
	}
}
