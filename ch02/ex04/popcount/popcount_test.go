package popcount_test

import (
	"popcount"
	"testing"
)

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountShift(0x1234567890ABCDEF)
	}
}

// goos: darwin
// goarch: arm64
// pkg: popcount
// BenchmarkPopCount-8             1000000000               0.3270 ns/op
// BenchmarkPopCountShift-8        53196009                22.17 ns/op
