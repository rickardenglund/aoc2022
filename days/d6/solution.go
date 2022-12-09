package d6

import (
	_ "embed"
	"fmt"
)

func p1(input string) string {
	return fmt.Sprintf("%d", findMarker([]rune(input), 4))
}

func p2(input string) string {
	return fmt.Sprintf("%d", findMarker([]rune(input), 14))
}

func findMarker(signal []rune, markerSize int) int {
	for i := range signal {
		m := map[rune]bool{}
		for n := 0; n < markerSize; n++ {
			m[signal[i+n]] = true
		}

		if len(m) == markerSize {
			return i + markerSize
		}
	}

	return -1
}
