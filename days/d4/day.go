package d4

import (
	_ "embed"
	"fmt"
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
	return fmt.Sprintf("%d", p1(input))
}

func (d Day) P2(input string) string {
	return fmt.Sprintf("%d", p2(input))
}

func (d Day) GetAnswer1() string {
	return "556"
}

func (d Day) GetAnswer2() string {
	return "876"
}
