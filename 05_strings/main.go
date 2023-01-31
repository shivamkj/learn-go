package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// STRING LITERALS VS RAW STRING LITERALS
	literal1 := "Hello, Go! stored in C:\\Users\\Documents\\go \n" // string literal: single line, interpreted
	literal2 :=
		`<h1>Hello<h1>,
\n \n	<h2>Go!<h2> Path: C:\Users\Documents\go` // raw string literal: multiline, not interpreted (no escape sequence processing)
	fmt.Println(literal1, literal2)
	// raw string literals are useful for writing HTML, XML, JSON, SQL, etc.

	// STRING LENGTH
	// len (built in function, no need for any import) returns the number of bytes in a string, works fine for ASCII
	// but not for UTF-8, use utf8.RuneCountInString() instead
	unicdoeName := "शिवम"
	fmt.Println(len(literal1), len(unicdoeName), utf8.RuneCountInString(unicdoeName))

	// STRING CONCATENATION
	fmt.Println("Hello" + " " + "Go!")

	// ESCAPE SEQUENCES
	// \n: new line, \t: horizontal tab, \v: vertical tab
	// \': single quote, \": double quote, \\: backslash
	fmt.Println("\nString Literal: \t Hello, \"Go!\"")

	// STRING UTILITIES
	// string package contains many useful functions to work with strings
	fmt.Println("Contains: ", strings.Contains(literal1, "Go!"))
	fmt.Println("Index: ", strings.Index(literal1, "o"), "Last Index: ", strings.LastIndex(literal1, "o"))
	replaced := strings.Replace(literal1, "Go!", "Shivam", 1)
	fmt.Println("Replace: ", replaced)
	fmt.Println("Upper: ", strings.ToUpper(literal1))
}
