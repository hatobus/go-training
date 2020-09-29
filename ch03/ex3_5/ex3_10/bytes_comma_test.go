package comma

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBytesComma(t *testing.T) {
	type testData struct {
		input  string
		output string
	}

	testCases := map[string]testData{
		"3文字以下の時": {
			input:  "100",
			output: "100",
		},
		"コンマが1つ入る": {
			input:  "1000",
			output: "1,000",
		},
		"コンマが2つ入る": {
			input:  "1000000",
			output: "1,000,000",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			out := BytesComma(tc.input)

			if diff := cmp.Diff(tc.output, out); diff != "" {
				t.Fatalf("invalid output, diff %v", diff)
			}
		})
	}
}
