package main

import "fmt"

func logicOperators() {
	a := true
	b := true

	//and (&&) (TODOS deben ser verdaderos)
	fmt.Println(a && b)

	//or (||) (UNO debe ser verdadero)
	fmt.Println(a || b)

	//not (!) (NEGAR todo)
	fmt.Println(!a) //false
	fmt.Println(!b) // true
}
