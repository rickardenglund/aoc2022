package slices_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"aoc/slices"
)

func TestContains(t *testing.T) {
	require.True(t, slices.Contains([]int{1, 2, 3}, 2))
	require.False(t, slices.Contains([]int{1, 2, 3}, 5))

	require.False(t, slices.Contains([]rune{'a'}, rune('4')))
	require.True(t, slices.Contains([]rune{'4'}, rune('4')))
}

func TestCount(t *testing.T) {
	require.Equal(t, 2, slices.Count([]int{1, 2, 2}, 2))
}
