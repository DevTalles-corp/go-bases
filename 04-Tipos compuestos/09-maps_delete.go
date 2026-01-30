package main

import "fmt"

func mapsDelete() {
	my_map := map[string]int{
		"Hola":  1,
		"Mundo": 2,
	}
	fmt.Println(my_map)

	delete(my_map, "Hola") //Un elemento del map
	clear(my_map)          //Limpiar todos los elementos
	fmt.Println(my_map)
	fmt.Println(len(my_map))
}
