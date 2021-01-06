package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCharCount(t *testing.T) {
	type testData struct {
		input  []string
		output string
	}

	testCases := map[string]testData{
		"ASCIIのみ": {
			input:  []string{"a", "b", "c"},
			output: fmt.Sprintf("%v\n", map[string]int{"ASCII_Hex_Digit": 3, "Hex_Digit": 3, "Pattern_White_Space": 3, "White_Space": 3}),
		},
		"日本語のみ": {
			input:  []string{"あ", "い", "う"},
			output: fmt.Sprintf("%v\n", map[string]int{"Pattern_White_Space": 3, "White_Space": 3}),
		},
		"日本語とASCIIが混在している": {
			input:  []string{"あ", "a", "い", "i", "う", "u"},
			output: fmt.Sprintf("%v\n", map[string]int{"ASCII_Hex_Digit": 1, "Hex_Digit": 1, "Pattern_White_Space": 6, "Soft_Dotted": 1, "White_Space": 6}),
		},
		"ASCIIとemojiが混在している": {
			input:  []string{"考え中", "🤔", "わいわい", "🙌", "おけまる〜", "🙆‍♀️"},
			output: fmt.Sprintf("%v\n", map[string]int{"Dash": 1, "Ideographic": 2, "Join_Control": 1, "Other_Math": 1, "Pattern_Syntax": 2, "Pattern_White_Space": 6, "Unified_Ideograph": 2, "Variation_Selector": 1, "White_Space": 6}),
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			cmd := exec.Command("go", "run", "charCount.go")

			var stdout bytes.Buffer
			cmd.Stdout = &stdout

			stdin, err := cmd.StdinPipe()
			if err != nil {
				t.Fatal(err)
			}

			go func() {
				defer stdin.Close()
				for _, in := range tc.input {
					io.WriteString(stdin, in+"\n")
				}
			}()

			if err := cmd.Run(); err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(stdout.String(), tc.output); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
