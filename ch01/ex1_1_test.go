package main

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hatobus/go-training/pioutil"
)

func TestEx1_1echo(t *testing.T) {
	type testData struct {
		Args []string
	}

	testCases := map[string]testData {
		"3つの引数を伴う時のテスト": {
			Args: []string{"a", "b", "c"},
		},
		"0個の引数の場合": {
			Args: []string{},
		},
		"日本語の場合": {
			Args: []string{"𠮷野家", "で", "𩸽"},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){
			os.Args = tc.Args

			out := pioutil.OutputCapture(func() {
				main()
			})

			if diff := cmp.Diff(out, strings.Join(tc.Args, " ")); diff != "" {
				t.Fatalf("unexpected output, diff: %v", diff)
			}
		})
	}
}
