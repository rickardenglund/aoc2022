package d13

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func p1(input string) string {
	iSum := 0
	pairs := strings.Split(input, "\n\n")
	for i, p := range pairs {
		ls := strings.Split(p, "\n")
		p1 := parsePacket(ls[0])

		p2 := parsePacket(ls[1])
		if p1.cmp(p2) <= 0 {
			iSum += i + 1
		}
	}

	return fmt.Sprintf("%d", iSum)
}

func p2(input string) string {
	packets := []packet{}
	for _, l := range strings.Split(input, "\n") {
		if l == "" {
			continue
		}

		packets = append(packets, parsePacket(l))
	}
	d1 := parsePacket("[[2]]")
	d2 := parsePacket("[[6]]")
	packets = append(packets, d1, d2)

	sort.Slice(packets, func(i, j int) bool {
		return packets[i].cmp(packets[j]) < 0
	})

	prod := 1
	for i, p := range packets {
		if p.String() == d1.String() || p.String() == d2.String() {
			prod *= i + 1

		}
	}

	return fmt.Sprintf("%d", prod)
}

func parsePacket(s string) packet {
	if n, err := strconv.Atoi(s); err == nil {
		return packet{v: n}
	}
	if !(s[0] == '[' && s[len(s)-1] == ']') {
		panic("invalid list: " + s)
	}

	p := packet{parts: []packet{}}
	i := 1
	currentNumber := ""
	for i < len(s) {
		if s[i] == '[' {
			endBracketIndex := matching(s, i)
			p.parts = append(p.parts, parsePacket(s[i:endBracketIndex+1]))
			i = endBracketIndex + 1
			continue
		}

		if s[i] == ',' || s[i] == ']' {
			if currentNumber == "" {
				i++
				continue
			}

			n, err := strconv.Atoi(currentNumber)
			if err != nil {
				panic(err)
			}
			currentNumber = ""
			p.parts = append(p.parts, packet{v: n})
			i++
			continue
		}

		currentNumber += string(s[i])
		i++
	}

	return p
}

func matching(parts string, start int) int {
	l := 0
	s := parts[start+1:]
	for i, c := range s {
		if c == '[' {
			l++
		}
		if c == ']' {
			l--
		}

		if l < 0 {
			return start + i + 1
		}
	}

	panic("no match")
}
