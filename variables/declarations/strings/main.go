package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare string
//
//  1. Declare a string variable
//
//  2. Print that variable
//
// EXPECTED OUTPUT
//  ""
// ---------------------------------------------------------

func main() {
	// USE THE BELOW CODE
	// You'll learn about Printf later

	var name string
	fmt.Printf("s (%T): %q\n", name, name)

	// %T prints the type of the value
	// %q prints an empty string
}
