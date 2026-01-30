package main

import "fmt"

func mapsInit() {

	//Map literal
	// teams := map[string][]string{
	// 	"Fernando Herrera FC":    []string{"Fernando", "Melissa", "Andrea"},
	// 	"Los co-instructores FC": []string{"Mariangel", "Gaston", "Isaias"},
	// 	"Los instructores":       []string{"Gabriel", "Teddy", "Ricardo"},
	// }

	// fmt.Println(teams["Los instructores"][0])

	ages := make(map[int][]string, 10)
	fmt.Println(ages)
	// Todos los que no se puedan comparar es decir usar == o !=
}
