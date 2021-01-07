package ex11_7

import (
	"testing"

	intset "github.com/hatobus/go-training/ch11/ex11_2"
)

func BenchmarkMapIntSetHas100(b *testing.B) {
	benchHas(b, intset.NewMapIntSet(), 100)
}

func BenchmarkMapIntSetHas1000(b *testing.B) {
	benchHas(b, intset.NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetHas10000(b *testing.B) {
	benchHas(b, intset.NewMapIntSet(), 10000)
}

func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchAdd(b, intset.NewMapIntSet(), 100)
}

func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchAdd(b, intset.NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetAdd10000(b *testing.B) {
	benchAdd(b, intset.NewMapIntSet(), 10000)
}

func BenchmarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, intset.NewMapIntSet(), intset.NewMapIntSet(), 100)
}

func BenchmarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, intset.NewMapIntSet(), intset.NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetUnionWith10000(b *testing.B) {
	benchUnionWith(b, intset.NewMapIntSet(), intset.NewMapIntSet(), 10000)
}

func BenchmarkMapIntSetCopy100(b *testing.B) {
	benchCopy(b, intset.NewMapIntSet(), 100)
}

func BenchmarkMapIntSetCopy1000(b *testing.B) {
	benchCopy(b, intset.NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetCopy10000(b *testing.B) {
	benchCopy(b, intset.NewMapIntSet(), 10000)
}
