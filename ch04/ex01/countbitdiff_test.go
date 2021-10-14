package main

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestCountBitDiff(t *testing.T) {
	var tests = []struct {
		x    [32]uint8
		y    [32]uint8
		want int
	}{
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("x")), 0},
		{[32]uint8{31: 1}, [32]uint8{30: 1, 31: 1}, 1},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CountBitDiff(%q, %q)", test.x, test.y)
		got := CountBitDiff(test.x, test.y)
		if got != test.want {
			t.Errorf("%s = %d want %d", descr, got, test.want)
		}
	}
}
