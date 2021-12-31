package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println(http.Get("ilmjtcv.com"))

	type Person struct {
		name string
		age  int
	}

	var gilad Person

	gilad.name = "Gilad"
	gilad.age = 20

	fmt.Println(json.Marshal(gilad))

	var talia Person

	fmt.Println(json.Unmarshal([]byte(`{age: 20, name: "Talia"}`), &talia))

}
