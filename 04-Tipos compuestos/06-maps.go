package main

import "fmt"

func mapsExample() {
	// [keyType]ValueType

	// var nilMap map[string]int
	// nilMap["Hola"] = 1 // Panic

	totalWins := map[string]int{}
	// fmt.Println(totalWins == nil) // false
	totalWins["Los domadores"] = 1
	totalWins["Los sedentarios"] = 2
	totalWins["Fernando Herrara FC"] = 5
	fmt.Println(totalWins["Fernando Herrara FC"])

}
