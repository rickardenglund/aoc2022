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

func Max[T Number](ns []T) T {
	var max T = ns[0]

	for _, n := range ns {
		if n > max {
			max = n
		}
	}

	return max
}
