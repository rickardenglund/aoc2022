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
)

func main() {
	days := []Day{d1.New(), d2.New(), d3.New()}

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

		tw.Append(toRow(name, p1Res, d.GetAnswer1(), p1Duration, p2Res, d.GetAnswer2(), p2Duration))
	}

	tw.Render()
}

func toRow(name string, p1Res, p1Expected int, p1Duration time.Duration, p2Res, p2Expected int, p2Dur time.Duration) []string {
	return []string{
		strings.TrimPrefix(name, "aoc/days/"),
		fmt.Sprintf("%v", toStatus(p1Res, p1Expected)),
		fmt.Sprintf("%v", toStatus(p2Res, p2Expected)),
		fmt.Sprintf("%v", p1Duration),
		fmt.Sprintf("%v", p2Dur),
	}
}

func toStatus(got int, expected int) string {
	if got == expected {
		return "success"
	}

	return fmt.Sprintf("expected %d got %d", expected, got)
}

type Day interface {
	P1(string) int
	P2(string) int
	GetInput() string
	GetAnswer1() int

	GetAnswer2() int
}
