package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(sli []int, index int) {
	temp := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = temp
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}

}

func main() {
	// var num_string string
	nums := make([]int, 0, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	num_string, _ := reader.ReadString('\n')
	num_string = strings.Replace(num_string, "\n", "", -1)
	var numlist = strings.Split(num_string, " ")
	for _, i := range numlist {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Invalid input! Please try again...")
			fmt.Printf("Please enter a list of integers (space-separated): ")
			continue
		}
		nums = append(nums, j)
	}
	BubbleSort(nums)
	fmt.Print(nums)
}
