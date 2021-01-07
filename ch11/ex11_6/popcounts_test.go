package ex11_6

import (
	"testing"

	"github.com/hatobus/go-training/ch02/ex2_4"
	"github.com/hatobus/go-training/ch02/ex2_5"
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
	doBench(b, ex2_4.PopCount)
}

func BenchmarkPopCountBitmagic(b *testing.B) {
	doBench(b, ex2_5.PopCount)
}
