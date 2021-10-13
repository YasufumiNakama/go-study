package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"+123456", "+123,456"},
		{"-123456", "-123,456"},
		{"+1234.56789", "+1,234.56789"},
		{"-1234.56789", "-1,234.56789"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("comma(%q)", test.s)
		got := comma(test.s)
		if got != test.want {
			t.Errorf("%s = %q want %q", descr, got, test.want)
		}
	}
}
