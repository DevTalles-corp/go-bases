package main

import "fmt"

const AppName = "Curso Go 2.0" //app = interno, App= Exportable
const MaxUsers = 1000

func constants() {
	var name string = "Ricardo"
	last_name := "Cuéllar"

	fmt.Printf("Nombre: %s %s %T\n", name, last_name, last_name)
	fmt.Println(AppName, MaxUsers)
}
