package slices

func Reverse[T any](input []T) []T {
	reversed := make([]T, len(input))
	for i := range input {
		reversed[len(reversed)-i-1] = input[i]
	}

	return reversed
}

func Rotate[T any](input [][]T) [][]T {
	res := make([][]T, 0, len(input))
	for x := range input[0] {
		col := []T{}
		for y := range input {
			col = append(col, input[y][x])
		}
		res = append(res, col)
	}

	return res
}
