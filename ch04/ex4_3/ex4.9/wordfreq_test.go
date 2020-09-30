package main

import (
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWordFreq(t *testing.T) {
	type testData struct {
		filename string
		output   string
	}

	testCases := map[string]testData{
		"Buffaloのテスト": {
			filename: "./testdata/buffalo.txt",
			output:   "map[Buffalo:3 buffalo:5]\n",
		},
		"thatのテスト": {
			filename: "./testdata/that.txt",
			output:   "map[He:1 boy:1 in:1 said:1 sentence:1 that:5 the:1 used:1 was:1 wrong:1]\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			args := []string{"run", "wordfreq.go", "-file", tc.filename}

			out, err := exec.Command("go", args...).Output()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(out), tc.output); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
