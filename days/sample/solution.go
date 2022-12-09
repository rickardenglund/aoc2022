package sample

import (
	_ "embed"
	"fmt"

	"aoc/parse"
)

func p1(input string) string {
	numbers := parse.Ints(input)

	return fmt.Sprintf("%d", numbers[0])
}

func p2(input string) string {
	numbers := parse.Ints(input)

	return fmt.Sprintf("%d", numbers[1])
}
