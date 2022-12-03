package main

import (
	_ "embed"
	"fmt"
	"reflect"
	"time"

	"aoc/days/d1"
	"aoc/days/d2"
	"aoc/days/d3"
	"aoc/days/sample"
)

func main() {
	days := []Day{sample.New(), d1.New(), d2.New(), d3.New()}

	for _, d := range days {
		name := reflect.TypeOf(d).PkgPath()
		fmt.Printf("%s:\n", name)

		start := time.Now()
		fmt.Printf("p1: %d - %v\n", d.P1(d.GetInput()), time.Since(start))
		start = time.Now()
		fmt.Printf("p2: %d - %v\n", d.P2(d.GetInput()), time.Since(start))
		fmt.Printf("\n")
	}
}

type Day interface {
	P1(string) int
	P2(string) int
	GetInput() string
}
