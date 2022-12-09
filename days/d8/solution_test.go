package d8

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/slices"
)

var testInput = `30373
25512
65332
33549
35390`

func Test_p1(t *testing.T) {
	require.Equal(t, "21", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer1(), d.P1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "8", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer2(), d.P2(puzzleInput))
}

func Test_look(t *testing.T) {
	require.Equal(t, 3, len(look([]tree{{0, 0}, {1, 1}, {2, 2}})))
	require.Equal(t, 2, len(look([]tree{{0, 0}, {1, 1}, {1, 2}})))
}

func Test_getScenicScore(t *testing.T) {
	hm := ReadHeightMap(testInput)
	r := slices.Rotate(hm)

	require.Equal(t, 4, getScenicScore(hm, r, 1, 2))
	require.Equal(t, 8, getScenicScore(hm, r, 3, 2))

}
