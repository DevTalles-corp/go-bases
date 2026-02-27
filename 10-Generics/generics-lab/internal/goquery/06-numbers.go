package goquery

// Sumar un slice de números
func Sum[T Number](in []T) T {
	var accumulate T
	for _, value := range in {
		accumulate += value
	}

	return accumulate
}
