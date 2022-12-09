package d8

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc/parse"
	"aoc/slices"
)

func p1(input string) string {
	heightMap := ReadHeightMap(input)
	rotatedMap := slices.Rotate(heightMap)
	seen := map[int]bool{}
	for _, r := range append(heightMap, rotatedMap...) {
		for _, t := range look(r) {
			seen[t] = true
		}

		revRow := slices.Reverse(r)
		for _, t := range look(revRow) {
			seen[t] = true
		}
	}

	return fmt.Sprintf("%d", len(seen))
}

func p2(input string) string {
	heightMap := ReadHeightMap(input)
	rotatedMap := slices.Rotate(heightMap)
	bestScore := 0
	for i := range heightMap {
		for j := range heightMap[0] {
			scenicScore := getScenicScore(heightMap, rotatedMap, i, j)
			if scenicScore > bestScore {
				bestScore = scenicScore
			}
		}
	}

	return fmt.Sprintf("%d", bestScore)
}

func getScenicScore(heightMap [][]tree, rotatedMap [][]tree, row int, col int) int {
	right := heightMap[row][col+1:]
	left := slices.Reverse(heightMap[row][:col])
	up := slices.Reverse(rotatedMap[col][:row])
	down := rotatedMap[col][row+1:]

	height := heightMap[row][col].height
	n := treeLook(height, right)
	n *= treeLook(height, left)
	n *= treeLook(height, up)
	n *= treeLook(height, down)

	return n
}

func treeLook(height int, trees []tree) int {
	n := 0
	for _, h := range trees {
		n++
		if h.height >= height {
			break
		}

	}

	return n
}

func look(trees []tree) []int {
	seen := []int{}
	maxHeight := -1
	for _, t := range trees {
		if t.height > maxHeight {
			seen = append(seen, t.pos)
			maxHeight = t.height
		}
	}

	return seen
}

type tree struct {
	height int
	pos    int
}

func (t tree) String() string {
	return fmt.Sprintf("%d", t.height)
}

func ReadHeightMap(input string) [][]tree {
	heightMap := [][]tree{}
	i := 0
	for _, r := range strings.Split(input, "\n") {
		row := []tree{}
		for _, h := range strings.Split(r, "") {
			t := tree{
				height: parse.Int(h),
				pos:    i,
			}
			i++
			row = append(row, t)
		}
		heightMap = append(heightMap, row)
	}

	return heightMap
}
