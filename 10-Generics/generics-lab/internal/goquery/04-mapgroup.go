package goquery

// Convertir []T a map[K]T
func IndexBy[T any, K comparable](items []T, keyFn func(T) K) map[K]T {
	out := make(map[K]T, len(items))
	for _, item := range items {
		out[keyFn(item)] = item
	}

	return out
}

// Agrupar elementos en map
func GroupBy[T any, K comparable](items []T, keyFn func(T) K) map[K][]T {
	out := make(map[K][]T)
	for _, item := range items {
		key := keyFn(item)
		out[key] = append(out[key], item)
	}

	return out
}
