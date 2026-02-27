package goquery

import "testing"

func TestReduceSum(test *testing.T) {
	nums := []int{1, 2, 3, 4}
	sum := Reduce(nums, 0, func(acc, number int) int { return acc + number })

	if sum != 10 {
		test.Fatalf("Queriamos 10 pero obtuvimos: %d", sum)
	}
}

func TestPipelineOrders(test *testing.T) {
	orders := []Order{
		{1, 10, 100},
		{2, 10, 200},
		{3, 20, 300},
		{4, 30, 400},
		{5, 10, 500},
	}

	customer10 := Filter(orders, func(order Order) bool { return order.Customer == 10 })
	totals := Map(customer10, func(order Order) int { return order.Total })
	sum := Reduce(totals, 0, func(accumulate, number int) int { return accumulate + number })

	if sum != 800 {
		test.Fatalf("Quería 300 y obtuvé %d", sum)
	}
}
