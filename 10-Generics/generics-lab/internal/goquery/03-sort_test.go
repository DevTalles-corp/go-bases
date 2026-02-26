package goquery

import "testing"

func TestSortAsc(test *testing.T) {
	in := []int{5, 1, 4, 2, 3}
	got := SortAsc(in)
	want := []int{1, 2, 3, 4, 5}

	for i := range want {
		if got[i] != want[i] {
			test.Fatalf("got[%d]=%v want=%v", i, got[i], want[i])
		}
	}
}

func TestSortByUsers(test *testing.T) {
	users := []User{{3, "Fernando"}, {2, "Beto"}, {1, "Ana"}}
	got := SortBy(users, func(a, b User) bool { return a.ID < b.ID })

	if got[0].ID != 1 || got[1].ID != 2 || got[2].ID != 3 {
		test.Fatalf("Orden no esperado: %+v", got)
	}
}
