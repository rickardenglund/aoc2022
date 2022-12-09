package d7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testInput = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func Test_p1(t *testing.T) {
	require.Equal(t, "95437", p1(testInput))
}

func Test_p1Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer1(), d.P1(puzzleInput))
}

func Test_p2(t *testing.T) {
	require.Equal(t, "24933642", p2(testInput))
}

func Test_p2Real(t *testing.T) {
	d := Day{}
	require.Equal(t, d.GetAnswer2(), d.P2(puzzleInput))
}
