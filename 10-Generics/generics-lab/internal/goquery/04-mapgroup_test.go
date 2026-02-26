package goquery

import "testing"

func TestIndexBy(test *testing.T) {
	users := []User{{1, "Ana"}, {2, "Beto"}, {3, "Fernando"}}
	idx := IndexBy(users, func(user User) int { return user.ID })

	if idx[1].Name != "Ana" || idx[2].Name != "Beto" || idx[3].Name != "Fernando" {
		test.Fatalf("Index no esperado: %+v", idx)
	}
}

type Order struct {
	ID       int
	Customer int
	Total    int
}

func TestGroupBy(test *testing.T) {
	orders := []Order{
		{1, 10, 100},
		{2, 10, 200},
		{3, 20, 300},
		{4, 30, 400},
	}

	got := GroupBy(orders, func(order Order) int { return order.Customer })

	if len(got[10]) != 2 || len(got[30]) != 1 {
		test.Fatalf("Grupo no esperado: %+v", got)
	}
}
