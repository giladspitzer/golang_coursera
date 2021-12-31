package main

func main() {
	type Person struct {
		name  string
		addr  string
		phone string
		age   int
	}

	var gilad Person

	gilad.name = "Gilad"
	gilad.age = 20

	talia := new(Person)

}
