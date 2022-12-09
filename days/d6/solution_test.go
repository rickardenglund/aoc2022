package d6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

func Test_p1(t *testing.T) {
	require.Equal(t, "7", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer1(), d.P1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "19", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer2(), d.P2(puzzleInput))
}

func Test_findMarker(t *testing.T) {
	tests := []struct {
		name   string
		signal []rune
		want   int
	}{
		{
			name:   "one",
			signal: []rune("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			want:   5,
		},
		{
			name:   "two",
			signal: []rune("nppdvjthqldpwncqszvftbrmjlhg"),
			want:   6,
		},
		{
			name:   "two",
			signal: []rune("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			want:   11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.signal, 4); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
