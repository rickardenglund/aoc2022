package main

import (
	"aoc/parse"
	"aoc/slices"
)

func p2(input string) int {
	return slices.Max(parse.Ints(input))
}
