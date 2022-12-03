package d3

import (
	_ "embed"
	"strings"

	"aoc/slices"
)

func p1(input string) int {
	sum := 0
	for _, l := range strings.Split(input, "\n") {
		compartment1 := l[:len(l)/2]
		compartment2 := l[len(l)/2:]

		doubles := inBoth([]rune(compartment1), []rune(compartment2))
		for _, l := range doubles {
			sum += getPrio(l)
		}
	}

	return sum
}

func getPrio(l rune) int {
	if l < 'a' { // capital
		return int(l) - 'A' + 27
	}

	// lowercase
	return int(l) - 'a' + 1
}

func inBoth(compartment1 []rune, compartment2 []rune) []rune {
	m := map[rune]bool{}
	for _, l := range compartment1 {
		m[l] = true
	}

	doubles := []rune{}
	for _, l := range compartment2 {
		if m[l] {
			doubles = append(doubles, l)
		}
	}

	return toUnique(doubles)
}

func toUnique(doubles []rune) []rune {
	m := map[rune]bool{}
	res := []rune{}
	for _, r := range doubles {
		m[r] = true
	}

	for r := range m {
		res = append(res, r)
	}

	return res
}

func p2(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for i := 0; i+3 <= len(lines); i += 3 {
		badge := findBadge(lines[i : i+3])
		sum += getPrio(badge)
	}

	return sum
}

func findBadge(backpacks []string) rune {
	for _, r := range backpacks[0] {
		if slices.Contains([]rune(backpacks[1]), r) &&
			slices.Contains([]rune(backpacks[2]), r) {
			return r

		}
	}

	panic("did not find badge")
}
