package main

import (
	"fmt"
	"testing"
)

func TestRmnextdup(t *testing.T) {
	var tests = []struct {
		array []string
		want  []string
	}{
		{[]string{"1", "1"}, []string{"1"}},
		{[]string{"1", "1", "1", "2", "2", "2"}, []string{"1", "2"}},
		{[]string{"1", "1", "2", "2", "1", "1", "3", "3", "3"}, []string{"1", "2", "1", "3"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("rmnextdup(%v)", test.array)
		got := rmnextdup(test.array)
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
