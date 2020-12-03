package popcount

import (
	"math/rand"
	"testing"
)

func BenchmarkPopCountOrigin(b *testing.B) {
	rand.Seed(12345)
	for i := 0; i < b.N; i++ {
		x := rand.Uint64()
		PopCountOrigin(x)
	}
}

func BenchmarkPopCountSyncOnce(b *testing.B) {
	rand.Seed(12345)
	for i := 0; i < b.N; i++ {
		x := rand.Uint64()
		PopCount(x)
	}
}
