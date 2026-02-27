package goquery

import "testing"

func TestPartition(test *testing.T) {
	// Pases un slice de numeros del 1 al 7 {1,2,3,4,5,6,7}
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// Obtener los dos slices ya separados
	even, odd := Partition(in, func(number int) bool { return number%2 == 0 })
	// Comparar los tamaños de cada slice que cumplan la condición

	if len(even) != 4 || len(odd) != 5 {
		test.Fatalf("No se esperaba par=%v impar=%v", even, odd)
	}
	// Dividir dos slices en numeros pares o impares
}
