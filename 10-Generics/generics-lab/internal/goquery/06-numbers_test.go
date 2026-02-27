package goquery

import "testing"

type MyInt int

func TestSum(test *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 {
		test.Fatal("Esperabamos 6")
	}

	if Sum([]float64{1.5, 2.5}) != 4.0 {
		test.Fatal("Esperabamos 4.0")
	}

	var x MyInt = 10
	var y MyInt = 20

	if Sum([]MyInt{x, y}) != 30 {
		test.Fatal("Esperabamos 30")
	}

}
