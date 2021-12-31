package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	type Name struct {
		fname string
		lname string
	}

	names := make([]Name, 0, 0)

	file, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		n := Name{fname: words[0], lname: words[1]}
		names = append(names, n)
		fmt.Println(names)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i].fname + " " + names[i].lname)
	}

}
