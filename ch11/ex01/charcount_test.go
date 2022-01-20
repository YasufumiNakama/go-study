package main

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	var tests = []struct {
		bytes   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			[]byte("ABCDE"),
			map[rune]int{'A': 1, 'B': 1, 'C': 1, 'D': 1, 'E': 1},
			[]int{0, 5, 0, 0, 0},
			0,
		},
		{
			[]byte("ABCDEあいうえお"),
			map[rune]int{'A': 1, 'B': 1, 'C': 1, 'D': 1, 'E': 1, 'あ': 1, 'い': 1, 'う': 1, 'え': 1, 'お': 1},
			[]int{0, 5, 0, 5, 0},
			0,
		},
		{
			[]byte{0xff},
			map[rune]int{},
			[]int{0, 0, 0, 0, 0},
			1,
		},
	}

	for _, test := range tests {
		counts, utflen, invalid, err := charcount(bytes.NewReader(test.bytes))
		if err != nil {
			t.Errorf("%v\n", err)
			continue
		}

		// counts 両側から検査
		match := true
		for k, v := range test.counts {
			if v != counts[k] {
				match = false
				break
			}
		}
		for k, v := range counts {
			if v != test.counts[k] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("[counts] got %v, want %v", counts, test.counts)
		}

		// utflen 両側から検査
		match = true
		for i, v := range test.utflen {
			if v != utflen[i] {
				match = false
				break
			}
		}
		for i, v := range utflen {
			if v != test.utflen[i] {
				match = false
				break
			}
		}
		if !match {
			t.Errorf("[utflen] got %v, want %v", utflen, test.utflen)
		}

		// invalid
		if invalid != test.invalid {
			t.Errorf("[invalid] got %v, want %v", invalid, test.invalid)
		}
	}
}
