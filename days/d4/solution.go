package d4

import (
	_ "embed"
	"strconv"
	"strings"
)

func p1(input string) int {
	count := 0
	sections := getSections(input)
	for _, s := range sections {
		if s.first.contains(s.other) || s.other.contains(s.first) {
			count++
		}
	}

	return count
}

type tuple struct {
	first, other section
}

func p2(input string) int {
	count := 0
	sections := getSections(input)
	for _, s := range sections {
		if s.first.overlaps(s.other) || s.other.overlaps(s.first) {
			count++
		}
	}

	return count
}

func getSections(input string) []tuple {
	var tuples []tuple
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}
		elves := strings.Split(l, ",")
		section1 := newSection(elves[0])
		section2 := newSection(elves[1])
		tuples = append(tuples, tuple{section1, section2})
	}

	return tuples
}

func newSection(elf string) section {
	numbers := strings.Split(elf, "-")
	min, err := strconv.Atoi(numbers[0])
	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(numbers[1])
	if err != nil {
		panic(err)
	}

	return section{
		min: min,
		max: max,
	}
}
