// composite data types hold lots of different types of data in one place

// Arrays
// • Fixed Length
// • Elements initialized to zero value of chosen type

package main

import (
	"fmt"
)

func main() {
	// var x [5]int = [5]{0, 1, 2, 3, 4}
	y := [...]int{0, 1, 2, 3, 4}
	fmt.Printf("%v", y)

	for i, v := range y {
		fmt.Printf("ind %d, val %d", i, v)
	}

	// slices are directly related to original array -- chaning one slice will change array and any other slices that access same parts that were changed
	a1 := y[0:2]
	fmt.Println(len(a1), cap(a1))

	sli := make([]int, 10)
	sli2 := make([]int, 10, 15) // type, length, capacity
	fmt.Printf("%v", sli)
	fmt.Printf("%v", sli2)

	sli3 := make([]int, 0, 3)
	sli3 = append(sli3, 100)
	fmt.Printf("%v", sli3)

}
