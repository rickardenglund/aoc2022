package d7

import (
	_ "embed"
	"fmt"
	"strings"

	"aoc/slices"
)

func p1(input string) string {
	const sizeLimit = 100000
	device := searchFileSystem(input)
	subDirs := device.root.allSubdirs()

	smallDirSizes := 0

	for _, sd := range subDirs {
		size := sd.getSize()
		if size <= sizeLimit {
			smallDirSizes += size
		}
	}

	return fmt.Sprintf("%d", smallDirSizes)
}

func p2(input string) string {
	const (
		totalDisk    = 70_000_000
		freeRequired = 30_000_000
	)

	device := searchFileSystem(input)
	totalFree := totalDisk - device.root.getSize()
	smallestToDelete := freeRequired - totalFree

	subDirs := device.root.allSubdirs()

	bigEnough := slices.Filter(subDirs, func(s *dir) bool {
		return s.getSize() > smallestToDelete
	})

	smallest := bigEnough[0].getSize()
	for _, d := range bigEnough {
		dSize := d.getSize()
		if dSize < smallest {
			smallest = dSize
		}
	}

	return fmt.Sprintf("%d", smallest)
}

func searchFileSystem(input string) device {
	d := device{
		root: dir{
			name:    "/",
			content: nil,
		},
	}
	d.cwd = &d.root

	li := 0
	lines := strings.Split(input, "\n")
	for li < len(lines) {
		if strings.HasPrefix(lines[li], "$") {
			cmd := strings.TrimPrefix(lines[li], "$ ")

			output := []string{}
			for li+1 < len(lines) && !strings.HasPrefix(lines[li+1], "$") {
				li++
				output = append(output, lines[li])
			}

			d.cmdExecuted(cmd, output)
		}

		li++
	}

	return d
}
