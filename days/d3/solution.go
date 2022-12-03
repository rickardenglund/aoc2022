package d3

import (
	_ "embed"
	"strings"

	"aoc/slices"
)

func p1(input string) int {
	sum := 0
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}
		compartment1 := l[:len(l)/2]
		compartment2 := l[len(l)/2:]

		l := inAll([]string{compartment1, compartment2})
		sum += getPrio(l)
	}

	return sum
}

func p2(input string) int {
	sum := 0
	lines := strings.Split(input, "\n")
	for i := 0; i+3 <= len(lines); i += 3 {
		badge := inAll(lines[i : i+3])
		sum += getPrio(badge)
	}

	return sum
}

func inAll(lists []string) rune {
	for _, l := range lists[0] {
		all := true
		for i := 1; i < len(lists); i++ {
			all = all && slices.Contains([]rune(lists[i]), l)
		}

		if all {
			return l
		}
	}

	panic("nothing found in all lists")
}

func getPrio(l rune) int {
	if l < 'a' { // capital
		return int(l) - 'A' + 27
	}

	// lowercase
	return int(l) - 'a' + 1
}
