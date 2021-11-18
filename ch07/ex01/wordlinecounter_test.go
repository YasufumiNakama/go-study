package main

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"hello world", 1},
		{"hello\nworld", 2},
		{"hello\nworld\n!", 3},
	}

	for _, test := range tests {
		var x LineCounter
		x.Write([]byte(test.s))
		if x != LineCounter(test.want) {
			t.Errorf("got %d want %d", x, test.want)
		}
	}
}

func TestWordCounter(t *testing.T) {
	var tests = []struct {
		s    string
		want int
	}{
		{"", 0},
		{"hello world", 2},
		{"hello\nworld\n!", 3},
	}

	for _, test := range tests {
		var x WordCounter
		x.Write([]byte(test.s))
		if x != WordCounter(test.want) {
			t.Errorf("got %d want %d", x, test.want)
		}
	}
}
