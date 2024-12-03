package main

import "fmt"

func main() {
	/*
		✅ A slice is basicly an array but, it can grow and shrink as you see fit
		➡️ Declaration:  slice_name := []datatype{values}
		🔗 Source: https://www.w3schools.com/go/go_slices.php
	*/

	slc := []int{23, 56, 75, 223}
	fmt.Printf("slc: %v\n", slc)

	// 💡 We can also start a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	slc2 := arr[1:4]

	fmt.Println("Capacidade do slc2:", slc2)
	fmt.Println("Capacidade do slc2:", cap(slc2))

	// 💡 The of changing elements values in a slice is the same of arrays
	prices := []float32{23.99, 45.99, 12.99, 128.99}
	prices[3] = 87.99
	fmt.Printf("price at index 3: $%v\n", prices[3])

	// 💡 To append an item to a slice, is quite simple
	// ➡️ slice_name = append(slice_name, element1, element2, ...)

	newPrices := []float32{6.99, 5.99, 4.99, 3.99, 2.99, 1.99}
	fmt.Printf("newPrices: %v\n", newPrices)

	prices = append(prices, 1.99, 3.99, 5.99) // manually add new items to 'prices'
	fmt.Printf("'prices' after appending new prices to it: %v\n", prices)

	// 💡 But we can also add all elements from an array to our slice like this
	brands := []string{"Puma", "Adidas", "Nike", "Mizuno"}
	newBrands := []string{"New Balance", "Hoka", "Asics"}

	fmt.Printf("'brands' BEFORE appending 'newBrands' elements to it: %v\n", brands)

	brands = append(brands, newBrands...) // 💡 using the spread operator to copy all the elements from one to another

	fmt.Printf("'brands' AFTER appending 'newBrands' elements to it: %v\n", brands)

	// 💡 We can change the length of a slice
	slicedBrands := brands[1:4] // ✅ Change length by re-slicing the array
	fmt.Printf("slicedBrands: %v\n", slicedBrands)
	fmt.Printf("length: %d\n", len(slicedBrands))
	fmt.Printf("capacity: %d\n", cap(slicedBrands))

	/*
		✅ Memory Efficiency
		💡 When using slices, Go loads all the underlying elements into the memory.
		   If the array is large and you need only a few elements,
		   it is better to copy those elements using the copy() function.
		➡️ Declaration: copy(dest, src)
	*/
	cars := []string{"Audi RS3", "Porsche Cayenne", "Mustang GT500"}

	fmt.Printf("Original cars list: %v\n", cars)

	neededCars := cars[:len(cars)-1]
	carsCopy := make([]string, len(neededCars))
	copy(carsCopy, neededCars)

	fmt.Printf("Copied cars 'carsCopy' list: %v\n", carsCopy)
}
