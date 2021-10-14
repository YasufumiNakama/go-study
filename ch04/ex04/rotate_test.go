package main

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		array []int
		k     int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 1, []int{4, 1, 2, 3}},
		{[]int{1, 2, 3, 4}, 2, []int{3, 4, 1, 2}},
		{[]int{1, 2, 3, 4}, 3, []int{2, 3, 4, 1}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 5, []int{4, 1, 2, 3}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("rotate(%v, %v)", test.array, test.k)
		got := rotate(test.array, test.k)
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
