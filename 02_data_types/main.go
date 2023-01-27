package main

import "fmt"

func main() {
	// ZERO VALUE
	// if variables are declared but not initialized, they are assigned the zero value
	// zero value for int is 0, for string is "", for bool is false, for float is 0.0
	// zero value for pointers, functions, interfaces, slices, channels, maps is nil
	// zero value for arrays & structs is the zero value of each field & element respectively
	// zero value for user defined types is the zero value of the underlying type

	//	Go Data Types: uint, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64,
	// complex64, complex128, bool, byte, rune, string, error, uintptr

	// integer types
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	fmt.Println(i, i8, i16, i32, i64)

	// float types (decimal numbers)
	var f32 float32
	var f64 float64
	fmt.Println(f32, f64)

	// booleans (TRUE or FALSE)
	var condition bool = true
	fmt.Print(condition)

	// string types
	var s string
	var r rune  // also a numeric type
	var by byte // also a numeric type
	fmt.Printf("%q\n", s)
	fmt.Println(r, by)

}
