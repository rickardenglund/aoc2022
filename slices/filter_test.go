package slices

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	require.Equal(t, []int{1, 2}, Filter([]int{1, 2, 3}, func(x int) bool { return x < 3 }))
}
