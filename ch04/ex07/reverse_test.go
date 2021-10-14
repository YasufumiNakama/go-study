package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		array []byte
		want  []byte
	}{
		{[]byte("こんにちは"), []byte("はちにんこ")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("reverse(%v)", string(test.array))
		reverse(test.array)
		match := true
		for i, v := range test.array {
			if v != test.want[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("%s = %v want %v", descr, string(test.array), string(test.want))
		}
	}
}
