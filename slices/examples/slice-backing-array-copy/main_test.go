package slice_backing_array_copy

import (
	"testing"
)

func generateTargetSliceFixture(length int) []string {
	res := make([]string, length)
	for i := 0; i < length; i++ {
		res[i] = "ice cream"
	}
	return res
}

const size = 50_000

func Benchmark_mapWithNilSlice(b *testing.B) {
	trg := generateTargetSliceFixture(size)
	for i := 0; i < b.N; i++ {
		mapWithNilSlice(trg)
	}
}

func Benchmark_mapWithPredefinedCapacity(b *testing.B) {
	trg := generateTargetSliceFixture(size)
	for i := 0; i < b.N; i++ {
		mapWithPredefinedCapacity(trg)
	}
}

func Benchmark_mapWithPredefinedLength(b *testing.B) {
	trg := generateTargetSliceFixture(size)
	for i := 0; i < b.N; i++ {
		mapWithPredefinedLength(trg)
	}
}
