package ex11_7

import (
	"math/rand"
	"testing"

	intset "github.com/hatobus/go-training/ch11/ex11_2"
)

const MAX_VALUE = 100000

func addRandomValue(set intset.IntSet, n int) {
	for i := 0; i < n; i++ {
		set.Add(i)
	}
}

func benchHas(b *testing.B, set intset.IntSet, n int) {
	addRandomValue(set, n)
	for i := 0; i < b.N; i++ {
		set.Has(rand.Intn(MAX_VALUE))
	}
}

func benchAdd(b *testing.B, set intset.IntSet, n int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(MAX_VALUE))
		}
		set.Clear()
	}
}

func benchUnionWith(b *testing.B, set1, set2 intset.IntSet, n int) {
	addRandomValue(set1, n)
	addRandomValue(set2, n)

	for i := 0; i < b.N; i++ {
		set1.UnionWith(set2)
	}
}

func benchCopy(b *testing.B, set intset.IntSet, n int) {
	addRandomValue(set, n)
	for i := 0; i < b.N; i++ {
		set.Copy()
	}
}

func BenchmarkWordIntSetHas100(b *testing.B) {
	benchHas(b, &intset.WordIntSet{}, 100)
}

func BenchmarkWordIntSetHas1000(b *testing.B) {
	benchHas(b, &intset.WordIntSet{}, 1000)
}

func BenchmarkWordIntSetHas10000(b *testing.B) {
	benchHas(b, &intset.WordIntSet{}, 10000)
}

func BenchmarkWordIntSetAdd100(b *testing.B) {
	benchAdd(b, &intset.WordIntSet{}, 100)
}

func BenchmarkWordIntSetAdd1000(b *testing.B) {
	benchAdd(b, &intset.WordIntSet{}, 1000)
}

func BenchmarkWordIntSetAdd10000(b *testing.B) {
	benchAdd(b, &intset.WordIntSet{}, 10000)
}

func BenchmarkWordIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, &intset.WordIntSet{}, &intset.WordIntSet{}, 100)
}

func BenchmarkWordIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, &intset.WordIntSet{}, &intset.WordIntSet{}, 1000)
}

func BenchmarkWordIntSetUnionWith10000(b *testing.B) {
	benchUnionWith(b, &intset.WordIntSet{}, &intset.WordIntSet{}, 10000)
}

func BenchmarkWordIntSetCopy100(b *testing.B) {
	benchCopy(b, &intset.WordIntSet{}, 100)
}

func BenchmarkWordIntSetCopy1000(b *testing.B) {
	benchCopy(b, &intset.WordIntSet{}, 1000)
}

func BenchmarkWordIntSetCopy10000(b *testing.B) {
	benchCopy(b, &intset.WordIntSet{}, 10000)
}
