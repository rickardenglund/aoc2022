package d3

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var testinput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestDay_P1(t *testing.T) {
	require.Equal(t, 157, p1(testinput))
}

func TestDay_P2(t *testing.T) {
	require.Equal(t, 70, p2(testinput))
}

func TestGetPrio(t *testing.T) {
	require.Equal(t, 1, getPrio('a'))
	require.Equal(t, 26, getPrio('z'))
	require.Equal(t, 27, getPrio('A'))
}

func TestFindBadge(t *testing.T) {
	bs := `fTTrrBqwfDTWfTDrRNrnRjgPSpJPnnmp
PvHPbsvZlMtbbvbCLLMHtHZZjtgJRjSnJSpSpjRgRjggSRmn
VLHbCbVPLZvlvMhHCHlPHbLCqQQfdQTBddTWhDTBchQzQwBW`
	lines := strings.Split(bs, "\n")

	require.Equal(t, 'P', findBadge(lines))
}
