package sample

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
	return len(strings.Split(input, "\n"))
}

func (d Day) P2(input string) int {
	return len(input)
}
