package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Printf("Please enter a string: ")
	fmt.Scan(&str)
	switch {
	case (strings.HasPrefix(str, "I") || strings.HasPrefix(str, "i")) && (strings.HasSuffix(str, "N") || strings.HasSuffix(str, "n")) && (strings.Contains(str, "A") || strings.Contains(str, "a")):
		fmt.Printf("Found!")
	default:
		fmt.Printf("Not Found!")
	}

}
