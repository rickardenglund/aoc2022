package d9

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func Test_p1(t *testing.T) {
	require.Equal(t, "13", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	require.Equal(t, "5902", p1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "1", p2(testInput))
}

var inputP2Large = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func Test_p2large(t *testing.T) {
	require.Equal(t, "36", p2(inputP2Large))
}
func Test_p2Real(t *testing.T) {
	require.Equal(t, "2445", p2(puzzleInput))
}

func Test_follow(t *testing.T) {
	tests := []struct {
		name    string
		head    coordinates
		tail    coordinates
		newTail coordinates
	}{
		{
			name:    "no step",
			head:    coordinates{1, 0},
			tail:    coordinates{0, 0},
			newTail: coordinates{0, 0},
		},
		{
			name:    "step right step",
			head:    coordinates{2, 0},
			tail:    coordinates{0, 0},
			newTail: coordinates{1, 0},
		},
		{
			name:    "step left step",
			head:    coordinates{-2, 0},
			tail:    coordinates{0, 0},
			newTail: coordinates{-1, 0},
		},
		{
			name:    "step diag",
			head:    coordinates{-2, 1},
			tail:    coordinates{0, 0},
			newTail: coordinates{-1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.newTail, follow(tt.head, tt.tail))
		})
	}
}

func Test_move(t *testing.T) {
	tests := []struct {
		name string
		rope rope
		dir  string
		want rope
	}{
		{
			name: "no step",
			rope: rope{{0, 0}, {0, 0}},
			dir:  "R",
			want: rope{{1, 0}, {0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rope.move(tt.dir)

			require.Equal(t, tt.want, tt.rope)
		})
	}
}
