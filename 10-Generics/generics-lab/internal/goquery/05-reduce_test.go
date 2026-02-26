package goquery

import "testing"

func TestReduceSum(test *testing.T) {
	nums := []int{1, 2, 3, 4}
	sum := Reduce(nums, 0, func(acc, number int) int { return acc + number })

	if sum != 10 {
		test.Fatalf("Queriamos 10 pero obtuvimos: %d", sum)
	}
}
