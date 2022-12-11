package sample

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc/parse"
)

func p1(input string) string {
	instructions := readInstructions(input)
	c := computer{x: 1, cycle: 0}

	nextCheck := 20
	signalSum := 0
	for _, inst := range instructions {
		newC := inst(c)
		if newC.cycle >= nextCheck {
			signalStrength := nextCheck * c.x
			signalSum += signalStrength
			nextCheck += 40
		}

		c = newC
	}

	return fmt.Sprintf("%d", signalSum)
}

func p2(input string) string {
	const (
		width  = 40
		height = 7
	)

	instructions := readInstructions(input)
	c := computer{x: 1, cycle: 0}
	screen := screen{
		pixels: make([]bool, width*height),
		width:  width,
		height: height,
	}

	currentCycle := 1
	for _, inst := range instructions {
		newC := inst(c)
		for currentCycle < newC.cycle {
			screen.draw(currentCycle, c.x)

			currentCycle++
		}

		c = newC
	}

	screen.pixels[0] = true
	return screen.String()
}

type screen struct {
	pixels        []bool
	width, height int
}

func (s screen) String() string {
	out := ""
	for i, p := range s.pixels {
		if i > 0 && i%s.width == 0 {
			out += "\n"
		}
		if p {
			out += "#"
		} else {
			out += "."
		}
	}

	return out
}

func shouldDraw(x, spritePos int) bool {
	d := spritePos - x

	return d >= -1 && d <= 1
}

func (s *screen) draw(cycle int, spritePos int) {
	if cycle >= len(s.pixels) {
		fmt.Printf("out of display: %d\n", cycle)
		return
	}

	s.pixels[cycle] = shouldDraw(cycle%s.width, spritePos)
}

type computer struct {
	x     int
	cycle int
}

func (c computer) String() string {
	return fmt.Sprintf("x: %d (%d)", c.x, c.cycle)
}

func readInstructions(input string) []func(c computer) computer {
	updates := []func(c computer) computer{}

	for _, l := range strings.Split(input, "\n") {
		parts := strings.Split(l, " ")
		switch parts[0] {
		case "addx":
			f := func(c computer) computer {
				n := parse.Int(parts[1])
				c.x += n
				c.cycle += 2

				return c
			}
			updates = append(updates, f)
		case "noop":
			f := func(c computer) computer {
				c.cycle += 1

				return c
			}
			updates = append(updates, f)
		}
	}

	return updates
}
