package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"aoc/slices"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Printf("p1: %d\n", p1(puzzleInput))
	fmt.Printf("p2: %d\n", p2(puzzleInput))
}

func p1(input string) int {
	elfCalories := getCalories(input)
	last := len(elfCalories) - 1

	return elfCalories[last]
}

func p2(input string) int {
	elfCalories := getCalories(input)
	last := len(elfCalories) - 1

	return slices.Sum(elfCalories[last-2:])
}

func getCalories(input string) []int {
	calories := 0
	var elfCalories []int

	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			elfCalories = append(elfCalories, calories)

			calories = 0
			continue
		}

		c, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		calories += c
	}

	sort.Ints(elfCalories)

	return elfCalories
}
