// This Program Demonstrates the Three Primary Formats for Declating and Instantiating Variables in Go

package main

import (
	"fmt"
)

var fruit string = "apple" // Compiler will care about unused variables in functions, but not at package level

func main() {
	var i float32 = 72 // This instantiation allows the user to define the type explicitly
	var j int          // declare but not instantiate
	j = 46
	f := 72 // In this declaration, the compiler determines the type (Note: Cannot be used at package level)

	{ // Block used for scoping to show variable shadowing
		var ( // Can group related variables for organization
			i = 42
			f = 72
		)
		fmt.Printf("%v, %T, %v, %T,%v, %T\n", i, i, f, f, j, j)
		// 42, int, 72, int, 46 int
	}
	fmt.Printf("%v, %T, %v, %T,%v, %T", i, i, f, f, j, j) // The %T indicates that the type of the value should be displayed
	// 72, float32, 72, int, 46, int

}
