package d7

import (
	"fmt"
)

type dir struct {
	name    string
	content []inode
	parent  *dir
}

func (d *dir) getName() string {
	return d.name
}

func (d *dir) getSize() int {
	sum := 0
	for _, c := range d.content {
		sum += c.getSize()
	}

	return sum
}

func (d *dir) getParent() *dir {
	return d.parent
}

func (d *dir) allSubdirs() []*dir {
	subDirs := []*dir{d}
	for _, subInode := range d.content {
		if subDir, ok := subInode.(*dir); ok {
			subDirs = append(subDirs, subDir.allSubdirs()...)
		}
	}

	return subDirs
}

func (d *dir) String() string {
	return fmt.Sprintf("Dir: %s", d.name)
}
