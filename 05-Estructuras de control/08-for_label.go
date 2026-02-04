package main

import "fmt"

func forLabel() {
	examples := []string{"Hola", "Carambola", "Bola", "Polola"}
outer:
	for _, example := range examples {
		for i, value := range example {
			fmt.Println(i, value, string(value))
			if value == 'l' {
				continue outer
			}
		}
	}
}
