package d8

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

func (d Day) P1(input string) string {
	return p1(input)
}

func (d Day) P2(input string) string {
	return p2(input)
}

func (d Day) GetAnswer1() string {
	return "1713"
}

func (d Day) GetAnswer2() string {
	return "268464"
}
