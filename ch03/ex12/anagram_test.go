package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	var tests = []struct {
		sa   string
		sb   string
		want bool
	}{
		{"", "", true},
		{"あいうえお", "うおあいえ", true},
		{"a b c", "  abc", true},
		{"abc", "abca", false},
		{"abca", "abc", false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("anagram(%q, %q)", test.sa, test.sb)
		got := anagram(test.sa, test.sb)
		if got != test.want {
			t.Errorf("%s = %t want %t", descr, got, test.want)
		}
	}
}
