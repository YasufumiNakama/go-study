package main

import (
	"testing"
)

func TestLen(t *testing.T) {
	var tests = []struct {
		vals []int
		want int
	}{
		{[]int{1}, 1},
		{[]int{1, 144, 9, 42}, 4},
	}

	for _, test := range tests {
		var x IntSet
		for _, v := range test.vals {
			x.Add(v)
		}
		got := x.Len()
		if got != test.want {
			t.Errorf("got %d want %d", got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		vals []int
		r    int
		want []int
	}{
		{[]int{1}, 1, []int{}},
		{[]int{1, 144, 9, 42}, 144, []int{1, 9, 42}},
	}

	for _, test := range tests {
		var x, want IntSet
		for _, v := range test.vals {
			x.Add(v)
		}
		for _, v := range test.want {
			want.Add(v)
		}
		x.Remove(test.r)
		if x.String() != want.String() {
			t.Errorf("got %s want %s", x.String(), want.String())
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		vals []int
		want []int
	}{
		{[]int{1}, []int{}},
		{[]int{1, 144, 9, 42}, []int{}},
	}

	for _, test := range tests {
		var x, want IntSet
		for _, v := range test.vals {
			x.Add(v)
		}
		for _, v := range test.want {
			want.Add(v)
		}
		x.Clear()
		if x.String() != want.String() {
			t.Errorf("got %s want %s", x.String(), want.String())
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		vals []int
		want []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 144, 9, 42}, []int{1, 144, 9, 42}},
	}

	for _, test := range tests {
		var x, want IntSet
		for _, v := range test.vals {
			x.Add(v)
		}
		for _, v := range test.want {
			want.Add(v)
		}
		got := x.Copy()
		if got.String() != want.String() {
			t.Errorf("got %s want %s", got.String(), want.String())
		}
	}
}
