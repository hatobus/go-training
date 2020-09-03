package ex2_4

import (
	"github.com/hatobus/go-training/ch02/ex2_3/popcount"
	"testing"
)

func TestPopCount(t *testing.T) {
	type testData struct {
		Input uint64
		ExpectOutput int
	}

	testCases := map[string]testData {
		"0の場合には0が出てくる": {
			Input: uint64(0), // 0
			ExpectOutput: int(0),
		},
		"7の場合には3が出てくる": {
			Input: uint64(7), // 111
			ExpectOutput: int(3),
		},
		"127の場合には7が出てくる": {
			Input: uint64(127), // 1111111
			ExpectOutput: int(7),
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			got := PopCount(tc.Input)

			if tc.ExpectOutput != got {
				t.Fatalf("PopCount unexpected output, expected %v, but got %v", tc.ExpectOutput, got)
			}
		})
	}
}

// goos: linux
// goarch: amd64
// pkg: github.com/hatobus/go-training/ch02/ex2_4
// BenchmarkPopCount
// BenchmarkPopCount-8           	23928091	        45.1 ns/op
// BenchmarkPopCountOriginal
// BenchmarkPopCountOriginal-8   	1000000000	         0.291 ns/op
// PASS

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(uint64(i))
	}
}
