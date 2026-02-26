package goquery

func Reduce[T any, R any](items []T, init R, fn func(accumulate R, item T) R) R {
	accumulate := init
	for _, item := range items {
		accumulate = fn(accumulate, item)
	}
	return accumulate
}
