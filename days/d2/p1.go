package d2

import (
	"fmt"
	"strings"
)

const (
	rock    move = 1
	paper   move = 2
	scissor move = 3
)

type move int

func round(me, other move) int {
	return int(me) + int(play(me, other))
}

const (
	loss outcome = 0
	draw outcome = 3
	win  outcome = 6
)

type outcome int

func play(me, other move) outcome {
	m := map[move]map[move]outcome{
		rock: {
			rock:    draw,
			paper:   loss,
			scissor: win,
		},
		paper: {
			rock:    win,
			paper:   draw,
			scissor: loss,
		},
		scissor: {
			rock:    loss,
			paper:   win,
			scissor: draw,
		},
	}

	return m[me][other]
}

func parseRow(r string) (move, move) {
	opponentMap := map[string]move{
		"A": rock,
		"B": paper,
		"C": scissor,
	}
	meMap := map[string]move{
		"X": rock,
		"Y": paper,
		"Z": scissor,
	}
	parts := strings.Split(r, " ")

	return meMap[parts[1]], opponentMap[parts[0]]
}

func parseRow2(r string) (desiredOutcome, move) {
	opponentMap := map[string]move{
		"A": rock,
		"B": paper,
		"C": scissor,
	}
	outComes := map[string]desiredOutcome{
		"X": needLoose,
		"Y": needDraw,
		"Z": needWin,
	}
	parts := strings.Split(r, " ")

	opponentMove := opponentMap[parts[0]]
	return outComes[parts[1]], opponentMove
}

type desiredOutcome int

const (
	needLoose desiredOutcome = iota
	needDraw
	needWin
)

func getMove(opponent move, outcome desiredOutcome) move {
	looseMap := map[move]move{
		rock:    scissor,
		paper:   rock,
		scissor: paper,
	}
	winMap := map[move]move{
		rock:    paper,
		paper:   scissor,
		scissor: rock,
	}

	switch outcome {
	case needLoose:
		return looseMap[opponent]
	case needDraw:
		return opponent
	case needWin:
		return winMap[opponent]
	default:
		panic(fmt.Sprintf("unknown outcome: %d", outcome))
	}
}
