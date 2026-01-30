package main

import "fmt"

func anonymousStruct() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Ricardo"
	person.age = 30
	person.pet = "Perro"

	fmt.Println(person.name)

	pet := struct {
		name string
		kind string
	}{
		name: "Dante",
		kind: "Perro",
	}

	fmt.Println(pet.kind)
}
