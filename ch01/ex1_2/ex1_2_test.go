package main

import (
	"github.com/google/go-cmp/cmp"
	"github.com/hatobus/go-training/pioutil"
	"os"
	"testing"
)

func TestEx2_2Echo(t *testing.T) {
	type testData struct {
		Args []string
		ExpectedOutput string
	}

	testCases := map[string]testData {
		"3つの引数を伴う場合のテスト": {
			Args: []string{"a", "b", "c"},
			ExpectedOutput: "0a 1b 2c",
		},
		"0個の引数の場合": {
			Args: []string{},
			ExpectedOutput: "",
		},
		"日本語の場合": {
			Args: []string{"𠮷野家", "で", "𩸽"},
			ExpectedOutput: "0𠮷野家 1で 2𩸽",
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			os.Args = tc.Args

			out := pioutil.OutputCapture(func() {
				main()
			})

			if diff := cmp.Diff(out, tc.ExpectedOutput); diff != "" {
				t.Fatalf("unexpected output, diff: %v", diff)
			}
		})
	}
}
