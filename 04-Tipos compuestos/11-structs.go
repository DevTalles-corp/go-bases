package main

import "fmt"

func structsExample() {
	type person struct {
		name string
		age  int
		pet  string
	}

	// var ricardo person

	// Struct litetal
	// fernando := person{}

	ricardo := person{
		"Ricardo",
		20,
		"Perro",
	}

	fernando := person{
		age:  40,
		pet:  "Perros",
		name: "Fernando",
	}

	ricardo.name = "Andrei Ricardo"
	fmt.Println(ricardo.name, fernando.name)

}
