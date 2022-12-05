package d5

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func p1(input string) string {
	moves, state := parse(input)

	for _, m := range moves {
		state.move9000(m)
	}

	return state.tops()
}

func p2(input string) string {
	moves, state := parse(input)

	for _, m := range moves {
		state.move9001(m)
	}

	return state.tops()
}

type move struct {
	n, from, to int
}

type state struct {
	cols [][]rune
}

func (s state) count() int {
	cnt := 0
	for _, c := range s.cols {
		cnt += len(c)
	}

	return cnt
}
func (s state) print() {
	fmt.Printf("n: %d\n", s.count())

	for _, c := range s.cols {
		fmt.Printf("%s\n", string(c))
	}
}
func (s *state) move9000(m move) {
	for i := 0; i < m.n; i++ {

		box := s.cols[m.from][len(s.cols[m.from])-1]
		s.cols[m.from] = s.cols[m.from][:len(s.cols[m.from])-1]
		s.cols[m.to] = append(s.cols[m.to], box)
	}
}

func (s *state) move9001(m move) {
	boxes := s.cols[m.from][len(s.cols[m.from])-m.n : len(s.cols[m.from])]
	if len(boxes) != m.n {
		panic("no")
	}
	s.cols[m.from] = s.cols[m.from][:len(s.cols[m.from])-m.n]
	s.cols[m.to] = append(s.cols[m.to], boxes...)
}

func (s state) tops() string {
	res := ""
	for _, c := range s.cols {
		res += string(c[len(c)-1])

	}

	return res
}

func parse(input string) ([]move, state) {
	parts := strings.Split(input, "\n\n")

	state := parseState(parts[0])
	moves := parseMoves(parts[1])

	return moves, state
}

func parseState(stateString string) state {
	rows := strings.Split(stateString, "\n")

	indexRow := rows[len(rows)-1]
	var cols []int
	for i := 1; true; i++ {
		index := strings.Index(indexRow, fmt.Sprintf("%d", i))
		if index == -1 {
			break
		}

		cols = append(cols, index)

	}

	s := make([][]rune, len(cols))
	for r := len(rows) - 2; r >= 0; r-- {
		for ci, cpos := range cols {
			if cpos < len(rows[r]) && rows[r][cpos] != ' ' {
				s[ci] = append(s[ci], rune(rows[r][cpos]))
			}

		}
	}

	return state{cols: s}
}

func parseMoves(movesString string) []move {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	var moves []move

	for _, mString := range strings.Split(movesString, "\n") {
		matches := re.FindStringSubmatch(mString)
		n, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		from, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		to, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		moves = append(moves, move{
			n:    n,
			from: from - 1,
			to:   to - 1,
		})
	}
	return moves
}
