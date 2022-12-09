package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotate(t *testing.T) {
	require.Equal(t, [][]int{{1, 3}, {2, 4}}, Rotate([][]int{{1, 2}, {3, 4}}))
	require.Equal(t, [][]int{{1, 4}, {2, 5}, {3, 6}}, Rotate([][]int{{1, 2, 3}, {4, 5, 6}}))
}
