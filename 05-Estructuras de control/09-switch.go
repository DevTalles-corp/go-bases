package main

import "fmt"

func switchExample() {
loop:
	for i := 0; i < 10; i++ {
		// switch size := len(word); size{}
		switch i {
		case 0, 2, 4, 6, 8:
			fmt.Println(i, " es par")
		case 3:
			fmt.Println(i, "es divisible entre 3")
		case 7:
			fmt.Println("Aqui termina el programa")
			break loop
		default:
			fmt.Println(i, "Valor por defecto")
		}
	}
}
