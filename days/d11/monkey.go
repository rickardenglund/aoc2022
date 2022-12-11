package d11

import (
	"fmt"
	"regexp"
	"strings"

	"aoc/parse"
)

type worryLevel int

type monkey struct {
	items      []worryLevel
	operation  func(i worryLevel) worryLevel
	test       func(i worryLevel) int
	nInspected int
	divisor    int
}

func (m monkey) String() string {
	return fmt.Sprintf("n: %d, %v", m.nInspected, m.items)
}

func parseMonkey(mStr string) monkey {
	re := regexp.MustCompile(`Monkey \d+:
  Starting items: ([\d, ]+)
  Operation: new = (old [*+] [\dold]+)
  Test: divisible by (\d+)
    If true: throw to monkey (\d+)
    If false: throw to monkey (\d)`)

	groups := re.FindStringSubmatch(mStr)
	if len(groups) != 6 {
		panic("failed to parse monkey: \n" + mStr)
	}

	return monkey{
		items:      parseWorryLevels(groups[1]),
		operation:  parseOperation(groups[2]),
		test:       parseTestOperation(groups[4], groups[5], groups[3]),
		divisor:    parse.Int(groups[3]),
		nInspected: 0,
	}
}

func parseTestOperation(ifTrueStr, ifFalseStr, divisorStr string) func(w worryLevel) int {
	ifTrue := parse.Int(ifTrueStr)
	ifFalse := parse.Int(ifFalseStr)

	a := func(w worryLevel) int {
		if w%worryLevel(parse.Int(divisorStr)) == 0 {
			return ifTrue
		} else {
			return ifFalse
		}
	}
	return a
}

func parseWorryLevels(worryLevelsString string) []worryLevel {
	ints := parse.IntList(worryLevelsString)
	levels := make([]worryLevel, len(ints))
	for i := range ints {
		levels[i] = worryLevel(ints[i])
	}

	return levels
}

func parseOperation(ops string) func(i worryLevel) worryLevel {
	parts := strings.Split(ops, " ")
	var op func(a, b worryLevel) worryLevel

	switch parts[1] {
	case "+":
		op = func(a, b worryLevel) worryLevel {
			return a + b
		}
	case "*":
		op = func(a, b worryLevel) worryLevel {
			return a * b
		}
	default:
		panic("unknown operation")
	}

	return func(w worryLevel) worryLevel {
		a1 := getArg(parts[0], w)
		a2 := getArg(parts[2], w)

		return op(a1, a2)
	}
}

func getArg(s string, w worryLevel) worryLevel {
	if s == "old" {
		return w
	}

	return worryLevel(parse.Int(s))
}
