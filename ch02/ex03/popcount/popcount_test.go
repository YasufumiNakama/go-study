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

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

// goos: darwin
// goarch: arm64
// pkg: popcount
// BenchmarkPopCount-8             1000000000               0.3266 ns/op
// BenchmarkPopCountLoop-8         309095090                3.891 ns/op
