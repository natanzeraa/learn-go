package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ✅ Simple Loops
	// 💡 Looping throught a list of strings
	arr := []string{"Duna", "How to get away with murder", "La Casa de Papel", "Supernatural", "Sandman"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 💡 Looping throught a fixed length integer
	for j := 0; j < 10; j++ {
		// 💡 if the condition is true, the loop will ignore the current index
		if j%2 != 1 {
			continue
		}
		fmt.Println(j)
	}

	// 💡 Looping throught a fixed length integer
	for k := 101; k < 111; k++ {
		// 💡 if the condition is true, the loop will stop at the current index
		if k%2 != 1 {
			break // 💡 in this case will stop at the first odd number found
		}
		fmt.Println(k)
	}

	// ✅ Nested Loops
	movies := []string{"Rambo", "Die Hard"}
	version := []int{1, 2, 3}

	ratedMovies := []string{}

	for m := 0; m < len(movies); m++ {
		for r := 0; r < len(version); r++ {
			ratedMovies = append(ratedMovies, movies[m]+" "+strconv.Itoa(version[r]))
		}
	}

	fmt.Println(ratedMovies)
}
