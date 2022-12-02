package d2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExample1(t *testing.T) {
	s := `A Y
B X
C Z`
	d := Day{}
	require.Equal(t, 15, d.P1(s))
}

func TestExample2(t *testing.T) {
	s := `A Y
B X
C Z`
	d := Day{}
	require.Equal(t, 12, d.P2(s))
}

func TestPlay(t *testing.T) {
	require.Equal(t, loss, play(paper, scissor))
	require.Equal(t, draw, play(scissor, scissor))
	require.Equal(t, win, play(rock, scissor))
}

func TestRound(t *testing.T) {
	require.Equal(t, 8, round(paper, rock))
}

func TestParseRow(t *testing.T) {
	me, other := parseRow("A Y")
	require.Equal(t, paper, me)
	require.Equal(t, rock, other)
}
