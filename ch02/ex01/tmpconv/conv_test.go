package tempconv

import (
	"testing"
)

func TestCToK(t *testing.T) {
	var tests = []struct {
		x    Celsius
		want Kelvin
	}{
		{AbsoluteZeroC, AbsoluteZeroK},
		{FreezingC, FreezingK},
		{BoilingC, BoilingK},
	}
	for _, test := range tests {
		got := CToK(test.x)
		if got != test.want {
			t.Errorf("CToK(%s) = %s want %s", test.x, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		x    Kelvin
		want Celsius
	}{
		{AbsoluteZeroK, AbsoluteZeroC},
		{FreezingK, FreezingC},
		{BoilingK, BoilingC},
	}
	for _, test := range tests {
		got := KToC(test.x)
		if got != test.want {
			t.Errorf("KToC(%s) = %s want %s", test.x, got, test.want)
		}
	}
}
