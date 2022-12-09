package d9

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc/parse"
)

func p1(input string) string {
	return run(input, 2)

}

func p2(input string) string {
	return run(input, 10)
}

func run(input string, ropeSize int) string {
	dirs, steps := parseMoves(input)

	r := rope{}
	for i := 0; i < ropeSize; i++ {
		r = append(r, coordinates{0, 0})
	}

	visited := map[coordinates]bool{
		r.tail(): true,
	}
	for i := range dirs {
		for n := 0; n < steps[i]; n++ {
			r.move(dirs[i])
			visited[r.tail()] = true
		}
	}

	return fmt.Sprintf("%d", len(visited))
}

func parseMoves(input string) ([]string, []int) {
	var dirs []string
	var steps []int

	for _, r := range strings.Split(input, "\n") {
		parts := strings.Split(r, " ")
		dirs = append(dirs, parts[0])
		steps = append(steps, parse.Int(parts[1]))
	}

	return dirs, steps
}

func follow(target coordinates, f coordinates) coordinates {
	dx, dy := target.dist(f)

	newF := f
	newF.x += unit(dx)
	newF.y += unit(dy)

	if newF != target {
		return newF
	}

	return f
}

func unit(dy int) int {
	switch {
	case dy > 0:
		return 1
	case dy < 0:
		return -1
	}

	return 0
}
