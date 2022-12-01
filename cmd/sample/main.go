package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var puzzleInput string

func main() {
	fmt.Printf("p1: %d\n", p1(puzzleInput))
	fmt.Printf("p2: %d\n", p2(puzzleInput))
}
