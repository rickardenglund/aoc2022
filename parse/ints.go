package parse

import (
	"strconv"
	"strings"
)

func Ints(input string) []int {
	ns := []int{}
	for _, l := range strings.Split(input, "\n") {

		ns = append(ns, Int(l))
	}

	return ns
}

func Int(input string) int {
	n, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}

	return n
}
