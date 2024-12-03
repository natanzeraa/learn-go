package main

import "fmt"

func main() {

	/*
		✅ Declare arrays using the default syntax

		➡️ var array_name = [length]datatype{values} // here length is defined
		➡️ var array_name = [...]datatype{values} // here length is inferred

		⬆️
		⬇️

		➡️ array_name := [length]datatype{values} // here length is defined
		➡️ array_name := [...]datatype{values} // here length is inferred

		---------------------------------------------------
		🔗 Source: https://www.w3schools.com/go/go_arrays.php
		---------------------------------------------------
	*/

	// Declaring arrays
	var arr1 = [3]string{"First", "Second", "Third"}
	arr2 := [4]int8{12, 45, 120, 19}

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)

	// Arrays with inferred lengths
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]string{"First", "Second", "Third", "Fourth"}
	fmt.Printf("arr3: %v\n", arr3)
	fmt.Printf("arr4: %v\n", arr4)

	// Accessing array values
	ages := [3]int{26, 45, 76}
	fmt.Printf("ages[0]: %v\n", ages[0])
	fmt.Printf("ages[1]: %v\n", ages[1])
	fmt.Printf("ages[2]: %v\n", ages[2])

	// Modifying an array element
	ages[2] = 12
	fmt.Printf("ages[2] after modification: %v\n", ages[2])

	// We can also fully, partialy or choose to not initialize an array, for example 👇
	arr5 := [5]int{}
	arr6 := [6]int{1, 2, 3}
	arr7 := [7]int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("arr5: %v\n", arr5)
	fmt.Printf("arr6: %v\n", arr6)
	fmt.Printf("arr7: %v\n", arr7)

	// We can partially initialize certain elements of our array, for example 👇
	arr8 := [8]int{1: 34, 7: 66}
	fmt.Printf("arr8: %v\n", arr8)

	// Finding the length of an array
	fmt.Printf("The of arr5 is: %v\n", len(arr5))
	fmt.Printf("The of arr7 is: %v\n", len(arr7))
}
