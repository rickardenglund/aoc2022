package d2

import (
	_ "embed"
	"fmt"
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

func (d Day) P1(input string) string {
	lines := strings.Split(input, "\n")
	score := 0
	for _, line := range lines {
		me, opponent := parseRow(line)
		score += round(me, opponent)
	}

	return fmt.Sprintf("%d", score)
}

func (d Day) P2(input string) string {
	lines := strings.Split(input, "\n")

	score := 0
	for _, line := range lines {
		desiredOutcome, opponent := parseRow2(line)
		myMove := getMove(opponent, desiredOutcome)
		score += round(myMove, opponent)
	}

	return fmt.Sprintf("%d", score)
}

func (d Day) GetAnswer1() string {
	return "11906"
}

func (d Day) GetAnswer2() string {
	return "11186"
}
