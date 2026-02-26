package goquery

import "testing"

func TestContainsAndIndexOf(test *testing.T) {
	in := []string{"z", "b", "c"}

	if !Contains(in, "b") {
		test.Fatal("Se esperaba una 'b'")
	}

	if Contains(in, "y") {
		test.Fatal("No se esperaba 'z'")
	}

	if IndexOf(in, "c") != 2 {
		test.Fatal("Esperabamos un índice 2")
	}

	if IndexOf(in, "z") != -1 {
		test.Fatal("Se esperaba -1")
	}

}

func TestUnique(test *testing.T) {
	in := []int{1, 1, 2, 3, 2, 4, 4}
	got := Unique(in)
	want := []int{1, 2, 3, 4}

	if len(got) != len(want) {
		test.Fatalf("len got=%d want=%d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			test.Fatalf("got[%d]=%v  want=%v", i, got[i], want[i])
		}
	}
}
