package strings

import (
	"testing"
)

func TestJoin(t *testing.T) {
	var tests = []struct {
		sep string
		elems []string
		want string
	}{
		{"", nil, ""},
		{" ", []string{"a"}, "a"},
		{" ", []string{"a", "b", "c"}, "a b c"},
		{" ", []string{"123", "456", "789"}, "123 456 789"},
	}

	for _, test := range tests {
		got := Join(test.sep, test.elems...)
		if got != test.want {
			t.Errorf("got %s want %s", got, test.want)
		}
	}
}
