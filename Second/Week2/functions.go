package main

import (
	"fmt"
	"math"
)

var funcVar func(int) int

func incFunc(x int) int {
	return x + 1
}

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func MakeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

// Environment of a function is the set of names tha are valid in the function
// Closure: When functions are passed/returned their correspondng environments come with them

func main() {
	funcVar = incFunc
	fmt.Print(funcVar(1))

	v := applyIt(func(x int) int { return x + 1 }, 2)
	fmt.Print(v)

	Dist1 := MakeDistOrigin(0, 0)
	Dist2 := MakeDistOrigin(2, 2)
	fmt.Println(Dist1(2, 2))
	fmt.Println(Dist2(2, 2))
}
