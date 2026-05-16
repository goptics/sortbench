package sortbench

import (
	"math/rand"
	"testing"
)

func genSlice(size int) []int {
	s := make([]int, size)
	for i := range s {
		s[i] = rand.Intn(size)
	}
	return s
}

func BenchmarkSort_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := genSlice(100)
		b.StartTimer()
		Sort(nums)
	}
}

func BenchmarkSort_1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := genSlice(1000)
		b.StartTimer()
		Sort(nums)
	}
}

func BenchmarkSort_10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := genSlice(10000)
		b.StartTimer()
		Sort(nums)
	}
}
