package goquery

func Partition[T any](items []T, fn func(T) bool) (yes []T, no []T) {
	// Creación dos slices vacios
	yes = make([]T, 0, len(items))
	no = make([]T, 0, len(items))
	for _, item := range items {
		if fn(item) {
			yes = append(yes, item)
		} else {
			no = append(no, item)
		}
	}

	return yes, no
}
