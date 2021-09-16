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

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountClear(0x1234567890ABCDEF)
	}
}

// goos: darwin
// goarch: arm64
// pkg: popcount
// BenchmarkPopCount-8             1000000000               0.3243 ns/op
// BenchmarkPopCountClear-8        100000000               11.65 ns/op
