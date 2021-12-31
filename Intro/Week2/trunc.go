package main

import "fmt"

func main() {
	var num float64
	fmt.Printf("Please enter a floating point number: ")
	fmt.Scan(&num)
	fmt.Println(int64(num))
}
