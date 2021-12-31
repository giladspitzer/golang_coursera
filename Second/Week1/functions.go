package main

import "fmt"

func foo(y *int) {
	*y =
		*y + 1
}

func bar(y int) {
	y = y + 2
}

func main() {
	x := 2
	k := 2
	bar(k)
	foo(&x)
	fmt.Print(x)
	fmt.Print(k)
}
