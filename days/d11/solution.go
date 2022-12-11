package d11

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

func p1(input string) string {
	return simulate(input, 20, true)
}

func p2(input string) string {
	return simulate(input, 10_000, false)
}

func simulate(input string, rounds int, applyWorryrelief bool) string {
	monkeys := parseMonkeys(input)

	prodAllMonkeyDivisors := 1
	for _, m := range monkeys {
		prodAllMonkeyDivisors *= m.divisor
	}

	for round := 1; round <= rounds; round++ {
		for mi := range monkeys {
			for _, item := range monkeys[mi].items {
				monkeys[mi].nInspected++

				newWorryLevel := monkeys[mi].operation(item)
				if newWorryLevel < 0 {
					panic("negative worry level")
				}

				if applyWorryrelief {
					newWorryLevel = newWorryLevel / 3
				}

				newWorryLevel = newWorryLevel % worryLevel(prodAllMonkeyDivisors)

				targetMonkey := monkeys[mi].test(newWorryLevel)
				monkeys[targetMonkey].items = append(monkeys[targetMonkey].items, newWorryLevel)
			}

			monkeys[mi].items = nil
		}
	}

	var nInspections []int
	for _, m := range monkeys {
		nInspections = append(nInspections, m.nInspected)
	}

	sort.Ints(nInspections)
	lastIndex := len(nInspections) - 1

	return fmt.Sprintf("%d", nInspections[lastIndex]*nInspections[lastIndex-1])
}

func printRound(round int, monkeys []monkey) {
	fmt.Printf("round: %d\n", round)
	for _, m := range monkeys {
		fmt.Printf("%v\n", m)
	}
	fmt.Printf("\n")
}

func parseMonkeys(input string) []monkey {
	var monkeys []monkey
	for _, mStr := range strings.Split(input, "\n\n") {
		m := parseMonkey(mStr)

		monkeys = append(monkeys, m)
	}

	return monkeys
}
