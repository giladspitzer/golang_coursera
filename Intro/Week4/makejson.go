package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	idMap := make(map[string]string)
	fmt.Printf("Please enter a name: ")
	fmt.Scan(&name)
	fmt.Printf("Please enter an address: ")
	fmt.Scan(&address)
	idMap["name"] = name
	idMap["address"] = address

	fmt.Println(json.Marshal(idMap))

}
