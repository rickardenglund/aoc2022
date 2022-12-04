package d4

type section struct {
	min, max int
}

func (s section) contains(o section) bool {
	return s.min <= o.min && s.max >= o.max
}

func (s section) overlaps(o section) bool {
	return (s.min >= o.min && s.min <= o.max) ||
		(s.max <= o.max && s.max >= o.min)
}
