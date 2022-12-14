package d12

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func Test_p1(t *testing.T) {
	require.Equal(t, "31", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer1(), d.P1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "29", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer2(), d.P2(puzzleInput))
}
