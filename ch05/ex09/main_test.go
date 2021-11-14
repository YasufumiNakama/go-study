package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	var tests = []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"", strings.ToUpper, ""},
		{"$foo $bar", strings.ToUpper, "FOO BAR"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("expand(%v, %p)", test.s, test.f)
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf("%s = %q want %q", descr, got, test.want)
		}
	}
}
