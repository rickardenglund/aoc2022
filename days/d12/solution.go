package d12

import (
	_ "embed"
	"fmt"
	"strings"
)

func p1(input string) string {
	places, start, end := readMap(input)
	distance, err := findDistance(places, start, end)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%d", distance)
}

func p2(input string) string {
	places, _, end := readMap(input)
	starts := readStarts(input)
	min := 10000
	for p := range starts {
		distance, err := findDistance(places, p, end)
		if err != nil {
			continue
		}
		if distance < min {
			min = distance
		}
	}

	return fmt.Sprintf("%d", min)
}

type pos struct {
	x, y int
}

func (p pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func readStarts(input string) map[pos]bool {
	heightMap := map[pos]bool{}
	for y, l := range strings.Split(input, "\n") {
		for x, h := range l {
			if h == 'S' || h == 'a' {
				heightMap[pos{x, y}] = true
			}
		}
	}

	return heightMap
}

func readMap(input string) (map[pos]map[pos]bool, pos, pos) {
	heightMap := map[pos]int{}
	var start, end pos
	for y, l := range strings.Split(input, "\n") {
		for x, h := range l {
			if h == 'S' {
				h = 'a'
				start = pos{x, y}
			}
			if h == 'E' {
				h = 'z'
				end = pos{x, y}
			}

			heightMap[pos{x, y}] = int(h)
		}
	}

	//for y := 0; y < 8; y++ {
	//	for x := 0; x < 5; x++ {
	//		fmt.Printf("%c", heightMap[pos{x, y}])
	//	}
	//	fmt.Printf("\n")
	//}

	graph := map[pos]map[pos]bool{}
	for p, pv := range heightMap {
		possibleNeighbors := map[pos]bool{}
		for _, n := range neighbours(p) {
			if almost(heightMap[n], pv) {
				possibleNeighbors[n] = true
			}
		}
		graph[p] = possibleNeighbors
	}

	return graph, start, end
}

func almost(from int, to int) bool {
	//return from+1 <= to
	d := from - to

	return d <= 1
}

func neighbours(p pos) []pos {
	ns := []pos{
		//{p.x - 1, p.y - 1},
		{p.x, p.y - 1},
		//{p.x + 1, p.y - 1},

		{p.x - 1, p.y},
		{p.x + 1, p.y},

		//{p.x - 1, p.y + 1},
		{p.x, p.y + 1},
		//{p.x + 1, p.y + 1},
	}
	return ns
}
