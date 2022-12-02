package d2

import (
	"fmt"
	"strings"
)

const (
	rock    = 1
	paper   = 2
	scissor = 3
)

func round(me, other int) int {
	return me + play(me, other)
}

func play(me, other int) int {
	const (
		loss = 0
		even = 3
		win  = 6
	)

	switch me {
	case rock:
		switch other {
		case rock:
			return even
		case paper:
			return loss
		case scissor:
			return win
		default:
			panic(fmt.Sprintf("unknown move %d", other))
		}
	case paper:
		switch other {
		case rock:
			return win
		case paper:
			return even
		case scissor:
			return loss
		default:
			panic(fmt.Sprintf("unknown move %d", other))
		}
	case scissor:
		switch other {
		case rock:
			return loss
		case paper:
			return win
		case scissor:
			return even
		default:
			panic(fmt.Sprintf("unknown move %d", other))
		}
	default:
		panic(fmt.Sprintf("unknown me move %d", me))
	}
}

func parseRow(r string) (int, int) {
	opponentMap := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissor,
	}
	meMap := map[string]int{
		"X": rock,
		"Y": paper,
		"Z": scissor,
	}
	parts := strings.Split(r, " ")

	return meMap[parts[1]], opponentMap[parts[0]]
}

func parseRow2(r string) (int, int) {
	opponentMap := map[string]int{
		"A": rock,
		"B": paper,
		"C": scissor,
	}
	outComes := map[string]int{
		"X": needLoose,
		"Y": needDraw,
		"Z": needWin,
	}
	parts := strings.Split(r, " ")

	opponentMove := opponentMap[parts[0]]
	return getMove(opponentMove, outComes[parts[1]]), opponentMove
}

const (
	needLoose = iota
	needDraw
	needWin
)

func getMove(opponent int, outcome int) int {
	looseMap := map[int]int{
		rock:    scissor,
		paper:   rock,
		scissor: paper,
	}
	winMap := map[int]int{
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
