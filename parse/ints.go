package parse

import (
	"strconv"
	"strings"
)

func Ints(input string) []int {
	ns := []int{}
	for _, l := range strings.Split(input, "\n") {
		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		ns = append(ns, n)
	}

	return ns
}
