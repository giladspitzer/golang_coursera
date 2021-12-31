package main

import "fmt"

var x = 1 // heap memory
func main() {
	var x = 1 // stack memory
	func f(){
		fmt.Printf("%d", x)
	}
	func g(){
		fmt.Printf"%d", x)
	}

}
