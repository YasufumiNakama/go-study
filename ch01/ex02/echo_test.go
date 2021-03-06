// ref: https://github.com/adonovan/gopl.io/blob/master/ch11/echo/echo_test.go

package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"1", "2", "3"}, "0 1\n1 2\n2 3\n"},
		{[]string{"one", "two", "three"}, "0 one\n1 two\n2 three\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)

		out = new(bytes.Buffer)
		if err := echo(test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q want %q", descr, got, test.want)
		}
	}
}
