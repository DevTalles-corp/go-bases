package goquery

func Filter[T any](in []T, keep func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, value := range in {
		if keep(value) {
			out = append(out, value)
		}
	}
	return out
}

// Transformar []T a []R usando una función personalizada
func Map[T any, R any](in []T, fn func(T) R) []R {
	out := make([]R, 0, len(in))
	for _, value := range in {
		out = append(out, fn(value))
	}
	return out
}
