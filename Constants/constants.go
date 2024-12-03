package main

import "fmt"

// * CONSTANTS:
// ! Constant names are usually written in uppercase letters (for easy identification and differentiation from variables).
// ! Constants can be declared both inside and outside of a function.
// 🔗 Source: https://www.w3schools.com/go/go_constants.php

func main() {
	// * Defining constants
	const PI = 3.14 // ✅ This is an untyped constant, inferred automatically.
	fmt.Printf("\nPI: %v\n", PI)

	const SQUAREROOTOF16 int = 4 // ✅ This is a typed constant.
	fmt.Printf("\nSQUAREROOTOF16: %v\n", SQUAREROOTOF16)

	// ! Important:
	// ⚠️ Constants are immutable and cannot be changed during the application lifecycle.
	// If you try to modify a constant, like in the example below 👇, the compiler will throw an error:
	// 😵‍💫 cannot assign to A (neither addressable nor a map index expression).

	/*
		📌 Uncomment the lines below to see the compiler error:
		const A = 1
		A = 2
		fmt.Println(A)
	*/
}
