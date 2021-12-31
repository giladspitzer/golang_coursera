package main

import (
	"fmt"
)

func main() {
	var idMap map[string]int

	idMap = make(map[string]int)

	idMap2 := map[string]int{
		"Joe": 123,
	}

	fmt.Println(idMap2["Joe"])

	idMap2["Jane"] = 456

	delete(idMap2, "Joe")

	id, p := idMap2["Jane"]

	fmt.Println(id, p)

	fmt.Println(len(idMap2))

	for key, val := range idMap2 {
		fmt.Println(key, val)
	}

}
