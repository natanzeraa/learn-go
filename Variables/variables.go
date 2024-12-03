package main

import "fmt"

func main() {

	// * Variable declarations:
	// ✅ Full variable declaration with explicit type
	var number int8 = 12

	// ✅ Full variable declaration with inferred type
	var text = "I am a text . . ."

	// ✅ Short variable declaration, automatically inferring its type
	numberTwo := 2891

	fmt.Printf("\nnumber: %v\nnumberTwo: %v\ntext: %v\n", number, numberTwo, text)

	// * Variables with no initial values:
	// ! Variables in Go can be declared without an initial value, which results in their *zero value*:
	var a string // Zero value is an empty string
	var b int32  // Zero value is 0
	var c bool   // Zero value is false

	fmt.Printf("\na: %v\nb: %v\nc: %v\n", a, b, c)

	// * Assigning values after declaration:
	var myDog string      // Declared without an initial value
	myDog = "Barney Ross" // Assigned later
	fmt.Printf("\nmyDog: %v\n", myDog)

	/*
		📌 Differences between `var` and `:=` declarations:
		---------------------------------------------------
		var                                                 | :=
		---------------------------------------------------
		Can be used both inside and outside functions        | Can only be used inside functions
		Variable declaration and assignment can be separate | Declaration and assignment must occur together
		---------------------------------------------------
		🔗 Source: https://www.w3schools.com/go/go_variables.php
	*/

	// * Multiple variable declarations:
	// ✅ Declaring multiple variables with the same type
	var d, e, f, g = 20, 122, 77, 8
	fmt.Printf("\nd: %v\ne: %v\nf: %v\ng: %v\n", d, e, f, g)

	// ✅ Declaring multiple variables with different types
	var h, i, j = true, "Letter K", 34.90
	fmt.Printf("\nh: %v\ni: %v\nj: %v\n", h, i, j)

}
