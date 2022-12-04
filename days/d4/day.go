package d4

import (
	_ "embed"
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
	return p1(input)
}

func (d Day) P2(input string) int {
	return p2(input)
}

func (d Day) GetAnswer1() int {
	return 556
}

func (d Day) GetAnswer2() int {
	return 876
}
