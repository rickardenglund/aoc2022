package d2

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var puzzleInput string

func New() Day {
	return Day{}
}

type Day struct{}

func (d Day) GetInput() string {
	return puzzleInput
}

func (d Day) P1(input string) int {
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		me, opponent := parseRow(line)
		score += round(me, opponent)
	}

	return score
}

func (d Day) P2(input string) int {
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		me, opponent := parseRow2(line)
		score += round(me, opponent)
	}

	return score
}
