package main

import "fmt"

func inputOutput() {
	var name string
	var age int

	fmt.Println(name, age)
	fmt.Println(&name)
	fmt.Println(&age)
	fmt.Println("Ingresa tu nombre: ")
	fmt.Scan(&name)

	fmt.Println("Ingresa tu edad: ")
	fmt.Scan(&age)

	fmt.Printf("Hola %s, tienes %d años.\n", name, age)
	fmt.Println(&name)
	fmt.Println(&age)

}
