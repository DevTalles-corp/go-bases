package goquery

import "testing"

func TestFilterInts(test *testing.T) {
	in := []int{1, 2, 3, 4, 5}

	got := Filter(in, func(number int) bool { return number%2 == 0 })

	want := []int{2, 4}

	if len(got) != len(want) {
		test.Fatalf("len got=%d want=%d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			test.Fatalf("got[%d]=%v want=%v", i, got[i], want[i])
		}
	}
}

type User struct {
	ID   int
	Name string
}

func TestMapToNames(test *testing.T) {
	users := []User{{1, "Ana"}, {2, "Beto"}, {3, "Fernando"}}
	got := Map(users, func(user User) string { return user.Name })
	want := []string{"Ana", "Beto", "Fernando"}

	for i := range want {
		if got[i] != want[i] {
			test.Fatalf("got[%d] = %v want=%v", i, got[i], want[i])
		}
	}

}

func TestExplicitTypesArgsExample(test *testing.T) {
	users := []User{{1, "Ana"}, {2, "Beto"}, {3, "Fernando"}}

	got := Map[User, string](users, func(user User) string { return user.Name })

	want := []string{"Ana", "Beto", "Fernando"}

	for i := range want {
		if got[i] != want[i] {
			test.Fatalf("got[%d] = %v want=%v", i, got[i], want[i])
		}
	}
}
