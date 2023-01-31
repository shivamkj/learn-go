package main

import "fmt"

func main() {
	// ZERO VALUE
	// if variables are declared but not initialized, they are assigned the zero value
	// zero value for int is 0, for string is "", for bool is false, for float is 0.0
	// zero value for arrays & structs is the zero value of each field & element respectively
	// zero value for user defined types is the zero value of the underlying type

	// nil is a predeclared identifier representing the zero value for pointers, functions, interfaces, slices, channels, maps

	//	Go Data Types: uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64,
	// complex64, complex128, bool, byte, rune, string, error, uintptr

	// integer types
	var i int = 1234567890
	var i8 int8
	var i16 int16 = 6790
	var i32 int32 = 1234567890
	var i64 int64 = 1234567890123456789
	fmt.Println(i, i8, i16, i32, i64)

	// float types (decimal numbers)
	var f = 3.14
	var f32 float32
	var f64 float64
	fmt.Println(f, f32, f64)

	// booleans (TRUE or FALSE)
	var condition bool = true
	fmt.Print(condition)

	// string types
	var s string
	var r rune  // also a numeric type
	var by byte // also a numeric type
	fmt.Printf("%q\n", s)
	fmt.Println(r, by)

	// TYPE CONVERSION syntax: T(v) [type(value)]
	var distance int = 100
	var timer float64 = 25.5
	var speed float64 = float64(distance) / timer
	fmt.Println(speed)
	score := 100
	var increase = 5.5
	score = int(float64(score) * increase) // avoids loss of precision during type conversion
	fmt.Println(score)

}
