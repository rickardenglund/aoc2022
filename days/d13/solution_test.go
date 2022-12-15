package d13

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = "[1,1,3,1,1]\n[1,1,5,1,1]\n\n[[1],[2,3,4]]\n[[1],4]\n\n[9]\n[[8,7,6]]\n\n[[4,4],4,4]\n[[4,4],4,4,4]\n\n[7,7,7,7]\n[7,7,7]\n\n[]\n[3]\n\n[[[]]]\n[[]]\n\n[1,[2,[3,[4,[5,6,7]]]],8,9]\n[1,[2,[3,[4,[5,6,0]]]],8,9]\n"

func Test_p1(t *testing.T) {
	require.Equal(t, "13", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer1(), d.P1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "140", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer2(), d.P2(puzzleInput))
}

func TestParsePacket(t *testing.T) {
	tests := []struct {
		in string
	}{
		{in: "6"},
		{in: "[]"},
		{in: "[1]"},
		{in: "[1,2,3]"},
		{in: "[[1],[2,3,4]]"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("name; %d", i), func(t *testing.T) {
			packet := parsePacket(tt.in)

			require.Equal(t, tt.in, packet.String())
		})

	}
}

func TestMatching(t *testing.T) {
	tests := []struct {
		in    string
		start int
		want  int
	}{
		{in: "[1,[2,3]]", start: 3, want: 7},
		{in: "[[]]", start: 1, want: 2},
		{in: "[[1,[]]]", start: 1, want: 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			require.Equal(t, tt.want, matching(tt.in, tt.start))
		})
	}
}

func TestParsePkt2(t *testing.T) {
	for _, l := range strings.Split(testInput, "\n") {
		if l == "" {
			continue
		}

		require.Equal(t, l, parsePacket(l).String())
	}
}
