package slices

func Filter[T any](items []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(items))
	for _, item := range items {
		if predicate(item) {
			res = append(res, item)
		}
	}

	return res
}
