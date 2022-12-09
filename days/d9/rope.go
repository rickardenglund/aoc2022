package d9

import (
	"fmt"
)

type rope []coordinates

type coordinates struct {
	x, y int
}

func (r rope) String() string {
	s := ""
	for i := range r {
		s += fmt.Sprintf("%v", r[i])
	}

	return s
}

func (c coordinates) dist(o coordinates) (int, int) {
	return c.x - o.x, c.y - o.y
}

func (r rope) move(dir string) {
	switch dir {
	case "R":
		r[0].x += 1
	case "L":
		r[0].x -= 1
	case "U":
		r[0].y += 1
	case "D":
		r[0].y -= 1
	default:
		panic("unknown dir: " + dir)
	}

	for i := 0; i < len(r)-1; i++ {
		newpos := follow(r[i], r[i+1])
		r[i+1] = newpos
	}
}
func (r rope) tail() coordinates {
	return r[len(r)-1]
}
