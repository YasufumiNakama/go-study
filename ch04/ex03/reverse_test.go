package main

import (
	"fmt"
	"testing"
)

func TestCountBitDiff(t *testing.T) {
	var tests = []struct {
		array []int
		want  []int
	}{
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("reverse(%v)", test.array)
		reverse(&test.array)
		match := true
		for i, v := range test.array {
			if v != test.want[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("%s = %v want %v", descr, test.array, test.want)
		}
	}
}
