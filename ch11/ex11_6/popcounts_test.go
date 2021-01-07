package ex11_6

import (
	"testing"

	ex2_3 "github.com/hatobus/go-training/ch02/ex2_3/popcount"
)

func doBench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCountTable(b *testing.B) {
	doBench(b, PopCountTable)
}

func BenchmarkPopCountNaive(b *testing.B) {
	doBench(b, ex2_3.PopCount)
}
