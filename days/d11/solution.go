package d11

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"

	"aoc/parse"
)

func p1(input string) string {
	return simulate(input, 20, true)
}

func p2(input string) string {
	return simulate(input, 10_000, false)
}

func simulate(input string, rounds int, applyWorryrelief bool) string {
	monkeys := parseMonkeys(input)

	prodAllMonkeys := 1
	for _, m := range monkeys {
		prodAllMonkeys *= m.divisor
	}
	for round := 1; round <= rounds; round++ {
		for mi := range monkeys {
			for _, item := range monkeys[mi].items {
				monkeys[mi].nInspected++

				newWorryLevel := monkeys[mi].operation(item)
				if newWorryLevel < 0 {
					panic("invalid worrylevel")
				}
				if applyWorryrelief {
					newWorryLevel = newWorryLevel / 3
				}

				newWorryLevel = newWorryLevel % worryLevel(prodAllMonkeys)

				targetMonkey := monkeys[mi].test(newWorryLevel)

				monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, newWorryLevel)
			}

			monkeys[mi].items = nil
		}

		toDebug := map[int]bool{
			1:    true,
			20:   true,
			1000: true,
		}
		if toDebug[round] {
			printRound(round, monkeys)
		}
	}

	nInpsections := []int{}
	for _, m := range monkeys {
		nInpsections = append(nInpsections, m.nInspected)
	}

	sort.Ints(nInpsections)
	lastIndex := len(nInpsections) - 1

	return fmt.Sprintf("%d", nInpsections[lastIndex]*nInpsections[lastIndex-1])
}

func printRound(round int, monkeys []monkey) {
	fmt.Printf("round: %d\n", round)
	for _, m := range monkeys {
		fmt.Printf("%v\n", m)
	}
	fmt.Printf("\n")
}

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

func parseMonkeys(input string) []monkey {
	monkeys := []monkey{}
	for _, mStr := range strings.Split(input, "\n\n") {
		m := parseMonkey(mStr)

		monkeys = append(monkeys, m)
	}

	return monkeys
}

func parseMonkey(mStr string) monkey {
	m := monkey{}
	lines := strings.Split(mStr, "\n")
	for li := 0; li < len(lines); li++ {
		switch {
		case strings.HasPrefix(lines[li], "Monkey "), lines[li] == "":
		case strings.HasPrefix(lines[li], "  Starting items: "):
			ints := parse.IntList(strings.TrimPrefix(lines[li], "  Starting items: "))
			levels := make([]worryLevel, len(ints))
			for i := range ints {
				levels[i] = worryLevel(ints[i])
			}
			m.items = levels
		case strings.HasPrefix(lines[li], "  Operation: "):
			ops := strings.TrimPrefix(lines[li], "  Operation: new = ")
			m.operation = parseOperation(ops)
		case strings.HasPrefix(lines[li], "  Test: "):
			m.divisor = parse.Int(strings.TrimPrefix(lines[li], "  Test: divisible by "))
			li++
			ifTrue := parse.Int(strings.TrimPrefix(lines[li], "    If true: throw to monkey "))
			li++
			ifFalse := parse.Int(strings.TrimPrefix(lines[li], "    If false: throw to monkey "))

			m.test = func(w worryLevel) int {
				if w%worryLevel(m.divisor) == 0 {
					return ifTrue
				} else {
					return ifFalse
				}
			}
		default:
			panic("unknown row: " + lines[li])
		}
	}

	return m
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
