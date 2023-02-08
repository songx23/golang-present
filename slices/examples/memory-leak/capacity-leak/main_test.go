package main

import (
	"testing"
)

func Benchmark_getFirstFiveBytes(b *testing.B) {
	trg := make([]byte, 1_000_000)
	for i := 0; i < b.N; i++ {
		getFirstFiveBytes(trg)
	}
}
