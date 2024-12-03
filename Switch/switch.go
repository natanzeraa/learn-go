package main

import "fmt"

func main() {

	/*
		✅ Switch Case Statement
		💡 Use the switch statement to select one of many code blocks to be executed.
		➡️ Decalration:

			switch expression {
			case x:
			// code block
			case y:
			// code block
			case z:
			...
			default:
			// code block
			}
	*/

	// 💡 Let's find out if a number is divisible by 3, 5 or both
	randomNumber := 9

	switch {
	case randomNumber%3 == 0 && randomNumber%5 == 0:
		fmt.Printf("%v is divisible by 3 and also 5\n", randomNumber)
	case randomNumber%3 == 0:
		fmt.Printf("%v is divisible by 3\n", randomNumber)
	case randomNumber%5 == 0:
		fmt.Printf("%v is divisible by 5\n", randomNumber)
	default:
		fmt.Printf("%v is not divisible by 3 nor 5\n", randomNumber)
	}
}
