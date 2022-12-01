package slices

type Number interface {
	int | int64
}

func Sum[T Number](ns []T) T {
	var s T
	for _, n := range ns {
		s = s + n
	}

	return s
}
