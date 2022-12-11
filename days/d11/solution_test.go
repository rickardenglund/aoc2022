package d11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = "Monkey 0:\n  Starting items: 79, 98\n  Operation: new = old * 19\n  Test: divisible by 23\n    If true: throw to monkey 2\n    If false: throw to monkey 3\n\nMonkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0\n\nMonkey 2:\n  Starting items: 79, 60, 97\n  Operation: new = old * old\n  Test: divisible by 13\n    If true: throw to monkey 1\n    If false: throw to monkey 3\n\nMonkey 3:\n  Starting items: 74\n  Operation: new = old + 3\n  Test: divisible by 17\n    If true: throw to monkey 0\n    If false: throw to monkey 1\n"

func Test_p1(t *testing.T) {
	require.Equal(t, "10605", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer1(), d.P1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "2713310158", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer2(), d.P2(puzzleInput))
}

func Test_parseMonkey(t *testing.T) {
	monkeyString := "Monkey 1:\n  Starting items: 54, 65, 75, 74\n  Operation: new = old + 6\n  Test: divisible by 19\n    If true: throw to monkey 2\n    If false: throw to monkey 0"
	m := parseMonkey(monkeyString)
	require.Equal(t, []worryLevel{54, 65, 75, 74}, m.items)
	require.Equal(t, worryLevel(7), m.operation(1))
	require.Equal(t, 2, m.test(19))
	require.Equal(t, 0, m.test(20))
}
