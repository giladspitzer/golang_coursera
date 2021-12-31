package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int // ip is pointer to int

	ip = &x // value
	y = *ip // address pointer
	fmt.Println(ip, y)

	var k = 1 >> 5           // == 1 * 2 five times == 32
	var l = 32 << 5          // 32 / 2 five times == 1
	var strin = "helloWorld" // each string is comprised of characters which each represent a byte 8-bit, 16bit,32bit, etc. , each byte is a rune
	// strings are immutable
	// strconv pckg
	fmt.Println(k, l, strin)

	const (
		gilad = "Gilad"
		age   = 20
	)

	// using iota -- you dont care what the values are as long as they are all different
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		F
	)

	// for loop (and while)
	for x = 0; x < 5; x++ {
		fmt.Println(x)
	}
	// switch with tag
	switch x {
	case 3:
		fmt.Printf("Case 3")
	case 4:
		fmt.Printf("Case 4")
	case 5:
		fmt.Printf("Case 5")
	}

	// tagless switch
	switch {
	case x < 1:
		fmt.Printf("Case 1")
	case x > 1:
		fmt.Printf("Case 2")
	default:
		fmt.Printf("Case 3")
	}

	// break will exit loop, continue will end current iteration and continue next counter

	// scan
	var numApples int
	fmt.Printf("Number of apples?")
	// num, err :=
	fmt.Scan(&numApples)
	fmt.Println(numApples)

}
