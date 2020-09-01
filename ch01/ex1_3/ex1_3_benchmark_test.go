package main

import (
	"testing"
)

func generateStringSliceN(n int) []string {
	s := make([]string, n)

	for i := 0; i < n; i++ {
		s[i] = "a"
	}

	return s
}

// BenchmarkEchoFor
// BenchmarkEchoFor/100個の要素がある場合
// BenchmarkEchoFor/100個の要素がある場合-8                              	1000000000	         0.000014 ns/op
// BenchmarkEchoFor/1000個の要素がある場合
// BenchmarkEchoFor/1000個の要素がある場合-8                             	1000000000	         0.000375 ns/op
// BenchmarkEchoFor/10000個の要素がある場合
// BenchmarkEchoFor/10000個の要素がある場合-8                            	1000000000	         0.0272 ns/op
// BenchmarkEchoFor/100000個の要素がある場合
// BenchmarkEchoFor/100000個の要素がある場合-8                           	       1	2137555699 ns/op
// PASS
//
// Process finished with exit code 0

func BenchmarkEchoFor(b *testing.B) {
	type testData struct {
		Args []string
	}

	testCases := map[string]testData {
		"100個の要素がある場合": {
			Args: generateStringSliceN(100),
		},
		"1000個の要素がある場合": {
			Args: generateStringSliceN(1000),
		},
		"10000個の要素がある場合": {
			Args: generateStringSliceN(10000),
		},
		"100000個の要素がある場合": {
			Args: generateStringSliceN(100000),
		},
	}

	for testName, tc := range testCases {
		b.Run(testName, func(b *testing.B){
			b.ResetTimer()

			EchoFor(tc.Args)
		})
	}

}

// BenchmarkEchoJoin
// BenchmarkEchoJoin/100個の要素がある場合
// BenchmarkEchoJoin/100個の要素がある場合-8                           	1000000000	         0.000011 ns/op
// BenchmarkEchoJoin/1000個の要素がある場合
// BenchmarkEchoJoin/1000個の要素がある場合-8                          	1000000000	         0.000019 ns/op
// BenchmarkEchoJoin/10000個の要素がある場合
// BenchmarkEchoJoin/10000個の要素がある場合-8                         	1000000000	         0.000120 ns/op
// BenchmarkEchoJoin/100000個の要素がある場合
// BenchmarkEchoJoin/100000個の要素がある場合-8                        	1000000000	         0.00116 ns/op
// PASS
//
// Process finished with exit code 0

func BenchmarkEchoJoin(b *testing.B) {
	type testData struct {
		Args []string
	}

	testCases := map[string]testData {
		"100個の要素がある場合": {
			Args: generateStringSliceN(100),
		},
		"1000個の要素がある場合": {
			Args: generateStringSliceN(1000),
		},
		"10000個の要素がある場合": {
			Args: generateStringSliceN(10000),
		},
		"100000個の要素がある場合": {
			Args: generateStringSliceN(100000),
		},
	}

	for testName, tc := range testCases {
		b.Run(testName, func(b *testing.B){
			b.ResetTimer()

			EchoJoin(tc.Args)
		})
	}

}
