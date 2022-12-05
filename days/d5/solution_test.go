package d5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func Test_p1(t *testing.T) {
	require.Equal(t, "CMZ", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	require.Equal(t, "JCMHLVGMG", p1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "MCD", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	require.Equal(t, "LVMRWSSPZ", p2(puzzleInput))
}
