package convs

import (
	"testing"
)

func TestCToF(t *testing.T) {
	var tests = []struct {
		x    Celsius
		want Fahrenheit
	}{
		// {AbsoluteZeroC, AbsoluteZeroF}, // CToF(-273.15°C) = -459.66999999999996°F want -459.67°F
		{FreezingC, FreezingF},
		{BoilingC, BoilingF},
	}
	for _, test := range tests {
		got := CToF(test.x)
		if got != test.want {
			t.Errorf("CToF(%s) = %s want %s", test.x, got, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		x    Fahrenheit
		want Celsius
	}{
		{AbsoluteZeroF, AbsoluteZeroC},
		{FreezingF, FreezingC},
		{BoilingF, BoilingC},
	}
	for _, test := range tests {
		got := FToC(test.x)
		if got != test.want {
			t.Errorf("FToC(%s) = %s want %s", test.x, got, test.want)
		}
	}
}
