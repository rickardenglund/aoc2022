package main

import (
	_ "embed"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"

	"aoc/days/d1"
	"aoc/days/d2"
	"aoc/days/d3"
	"aoc/days/sample"
)

func main() {
	days := []Day{sample.New(), d1.New(), d2.New(), d3.New()}

	fmt.Printf("# Results\n")

	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader([]string{"day", "p1 result", "p2 result", "p1 time", "p2 time"})
	tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tw.SetCenterSeparator("|")

	for _, d := range days {
		name := reflect.TypeOf(d).PkgPath()

		start := time.Now()
		p1Res := d.P1(d.GetInput())
		p1Duration := time.Since(start)

		start = time.Now()
		p2Res := d.P2(d.GetInput())
		p2Duration := time.Since(start)

		tw.Append(toRow(name, p1Res, p1Duration, p2Res, p2Duration))
	}

	tw.Render()
}

func toRow(name string, p1Res int, p1Duration time.Duration, p2Res int, p2Dur time.Duration) []string {
	return []string{
		strings.TrimPrefix(name, "aoc/days/"),
		fmt.Sprintf("%d", p1Res),
		fmt.Sprintf("%d", p2Res),
		fmt.Sprintf("%v", p1Duration),
		fmt.Sprintf("%v", p2Dur),
	}
}

type Day interface {
	P1(string) int
	P2(string) int
	GetInput() string
}
