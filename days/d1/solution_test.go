package d1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func Test_p1(t *testing.T) {
	require.Equal(t, 24000, p1(testInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, 45000, p2(testInput))
}
