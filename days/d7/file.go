package d7

type file struct {
	name   string
	size   int
	parent dir
}

func (f *file) getName() string {
	return f.name
}

func (f *file) getSize() int {
	return f.size
}

func (f *file) getParent() *dir {
	return &f.parent
}
