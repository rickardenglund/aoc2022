package main

import (
	"aoc/parse"
	"aoc/slices"
)

func p1(input string) int {
	return slices.Sum(parse.Ints(input))
}
