package goquery

func Contains[T comparable](in []T, target T) bool {
	for _, value := range in {
		if value == target {
			return true
		}
	}
	return false
}

func Unique[T comparable](in []T) []T {
	seen := make(map[T]struct{}, len(in))
	out := make([]T, 0, len(in))

	for _, value := range in {
		if _, ok := seen[value]; ok {
			continue
		}
		seen[value] = struct{}{}
		out = append(out, value)
	}

	return out
}

// Devolver el indice del target enviado, sino esta o no existe regresa -1
func IndexOf[T comparable](in []T, target T) int {
	for index, value := range in {
		if value == target {
			return index
		}
	}

	return -1

}
