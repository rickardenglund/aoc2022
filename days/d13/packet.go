package d13

import (
	"fmt"
	"strings"
)

type packet struct {
	parts []packet
	v     int
}

func (p packet) String() string {
	if p.parts == nil {
		return fmt.Sprintf("%d", p.v)
	}

	parts := make([]string, len(p.parts))
	for i, p := range p.parts {
		parts[i] = p.String()
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ","))
}

func cmp(a, b int) int {
	return a - b
}

func (p packet) cmp(p2 packet) int {
	if p.parts == nil && p2.parts == nil {
		return cmp(p.v, p2.v)
	}

	if p.parts != nil && p2.parts != nil {
		for i := range p.parts {
			if i >= len(p2.parts) {
				break
			}
			d := p.parts[i].cmp(p2.parts[i])
			if d != 0 {
				return d
			}
		}
		if len(p.parts) < len(p2.parts) {
			return -1
		}
		if len(p2.parts) < len(p.parts) {
			return 1
		}

		return 0
	}

	if p.parts == nil {
		return packet{parts: []packet{p}}.cmp(p2)
	}

	if p2.parts == nil {
		return p.cmp(packet{parts: []packet{p2}})
	}

	panic("not impl")
}
