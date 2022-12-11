package d7

import (
	"strings"

	"aoc/parse"
)

type device struct {
	root dir
	cwd  *dir
}

func (d *device) cmdExecuted(cmd string, output []string) {
	parts := strings.Split(cmd, " ")
	switch parts[0] {
	case "cd":
		d.cd(parts[1])
	case "ls":
		d.ls(output)
	}
}

func (d *device) ls(outputs []string) {
	for _, line := range outputs {
		parts := strings.Split(line, " ")
		if parts[0] == "dir" {
			newDir := dir{
				name:    parts[1],
				content: nil,
				parent:  d.cwd}
			d.cwd.content = append(d.cwd.content, &newDir)
		} else {
			size := parse.Int(parts[0])
			newFile := file{
				name:   parts[1],
				size:   size,
				parent: *d.cwd,
			}
			d.cwd.content = append(d.cwd.content, &newFile)
		}
	}
}

func (d *device) cd(name string) {
	switch name {
	case "/":
		d.cwd = &d.root
	case "..":
		d.cwd = d.cwd.getParent()
	default:
		for _, subInode := range d.cwd.content {
			if subInode.getName() == name {
				if subDir, ok := subInode.(*dir); ok {
					d.cwd = subDir
				}

			}
		}

	}
}
