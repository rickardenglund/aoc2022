package d12

import (
	"fmt"
)

func findDistance(places map[pos]map[pos]bool, start pos, end pos) (int, error) {
	return findDistance_(places, map[pos]int{}, start, end)

}

func findDistance_(places map[pos]map[pos]bool, distances map[pos]int, start pos, end pos) (int, error) {
	toVisit := []pos{start}
	distances[start] = 0

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]

		if current == end {
			return distances[current], nil
		}

		currentDistance := distances[current]
		for n := range places[current] {
			if _, ok := distances[n]; !ok {
				distances[n] = currentDistance + 1
				toVisit = append(toVisit, n)
			}
		}
	}

	return 0, fmt.Errorf("no path found")
}

func dist(p pos, p2 pos) int {
	return abs(p.x-p2.x) + abs(p.y-p2.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
