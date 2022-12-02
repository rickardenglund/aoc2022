package d1

import (
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"aoc/slices"
)

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
