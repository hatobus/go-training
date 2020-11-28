package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkSVGPolygon(b *testing.B) {
	b.Run("naive", func(b *testing.B) {
		naive(ioutil.Discard)
	})

	var worker = 1
	// workerを2, 4, 8, 16, 32で増やしていく
	for i := 1; i <= 5; i++ {
		worker *= 2
		b.Run(fmt.Sprintf("Concurrent worker: %v", worker), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				concurrent(worker, ioutil.Discard)
			}
		})
	}
}
