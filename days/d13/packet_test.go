package d13

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_packet_Before(t *testing.T) {
	tests := []struct {
		packet  string
		packet2 string
		before  bool
	}{
		{
			packet:  "1",
			packet2: "1",
			before:  true,
		},
		{
			packet:  "[3]",
			packet2: "[1]",
			before:  false,
		},
		{
			packet:  "[3]",
			packet2: "[1]",
			before:  false,
		},
		{
			packet:  "[1,2]",
			packet2: "[1]",
			before:  false,
		},
		{
			packet:  "[1]",
			packet2: "[1,2]",
			before:  true,
		},
		{
			packet:  "[1]",
			packet2: "[[1,2],2]",
			before:  true,
		},
		{
			packet:  "[[1],[2,3,4]]",
			packet2: "[[1],4]",
			before:  true,
		},
		{
			packet:  "[[2,3,4]]",
			packet2: "[4]",
			before:  true,
		},
		{
			packet:  "[4,6,7]",
			packet2: "[4,6]",
			before:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.packet, func(t *testing.T) {
			p1 := parsePacket(tt.packet)
			p2 := parsePacket(tt.packet2)

			require.Equal(t, tt.before, p1.cmp(p2) <= 0)
		})
	}
}
