package d4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func Test_p1(t *testing.T) {
	require.Equal(t, 2, p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, "556", d.P1(d.GetInput()))

}

func Test_p2(t *testing.T) {
	require.Equal(t, 4, p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, "876", d.P2(d.GetInput()))

}
