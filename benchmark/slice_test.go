package benchmark

import (
	"github.com/hamidreza01/js.go/slice"
	"testing"
)

func BenchmarkTestPush(t *testing.B) {
	var s slice.Slice[int]
	for i := 0; i < t.N; i++ {
		s.Push(i)
	}
}

func BenchmarkTestMap(t *testing.B) {
	s := slice.Slice[int]{5, 5, 5, 5, 5}
	for i := 0; i < t.N; i++ {
		s.Map(func(_ int, v int) int {
			return v * 2
		})
	}
}

func BenchmarkTestFilter(t *testing.B) {
	s := slice.Slice[int]{5, 5, 1, 15, 100, 1000, 500, 650, 10, 5, 100}
	for i := 0; i < t.N; i++ {
		s.Filter(func(_ int, v int) bool {
			return v >= 10
		})
	}
}

func BenchmarkTestEvery(t *testing.B) {
	s := slice.Slice[int]{5, 5, 1, 15, 100, 1000, 500, 650, 10, 5, 100}
	for i := 0; i < t.N; i++ {
		s.Every(func(_ int, v int) bool {
			return v >= 10000
		})
	}
}
