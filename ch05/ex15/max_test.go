package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	var tests = []struct {
		vals []int
		valid bool
		want int
	}{
		{nil, false, 0},
		{[]int{1}, true, 1},
		{[]int{1, 2, 3, 4}, true, 4},
		{[]int{4, 3, 2, 1}, true, 4},
	}

	for _, test := range tests {
		got, err := max(test.vals...)
		if !test.valid {
			if err == nil {
				t.Errorf("There should be err != nil for %v\n", test.vals)
			}
			continue
		}
		if got != test.want {
			t.Errorf("got %d want %d", got, test.want)
		}
	}
}

func TestMaxAtLeastOneArg(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{4, 3, 2, 1}, 4},
	}

	for _, test := range tests {
		got := maxAtLeastOneArg(test.vals[0], test.vals[1:]...)
		if got != test.want {
			t.Errorf("got %d want %d", got, test.want)
		}
	}
}