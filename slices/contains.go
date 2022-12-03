package slices

func Contains[T comparable](xs []T, x T) bool {
	for _, c := range xs {
		if c == x {
			return true
		}
	}

	return false
}

func Count[T comparable](xs []T, x T) int {
	n := 0
	for _, c := range xs {
		if c == x {
			n++
		}
	}

	return n
}
