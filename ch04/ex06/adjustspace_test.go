package main

import (
	"fmt"
	"testing"
)

func TestAdjustspace(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte("xx   XX   xx"), []byte("xx XX xx")},
		{[]byte("  xx   XX   xx"), []byte(" xx XX xx")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("adjustspace(%v)", test.b)
		got := adjustspace(test.b)
		match := true
		for i, v := range got {
			if v != test.want[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("%s = %v want %v", descr, got, test.want)
		}
	}
}
