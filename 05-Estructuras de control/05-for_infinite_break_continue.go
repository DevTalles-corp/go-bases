package main

import "fmt"

func forInfinite() {
	n := 0
	for {
		n++

		if n == 100 {
			break
		}

		if n == 50 {
			fmt.Println("Continuar")
			continue
		}

		fmt.Println("Numero: ", n)
	}
}
