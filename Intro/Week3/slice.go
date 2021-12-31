package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s string
	var i int
	sli := make([]int, 0, 3)
	for {
		fmt.Printf("Please enter a number: ")
		_, err := fmt.Scan(&s)
		i, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Enter a valid number")
			if s == "X" {
				break
			} else {
				continue
			}
		} else {
			sli = append(sli, i)
			sort.Ints(sli[:])
		}
		fmt.Printf("%v", sli)
	}
}
