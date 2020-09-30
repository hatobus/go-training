package pointer_reverse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReversePointer(t *testing.T) {
	type testData struct {
		inputs [5]int
		output [5]int
	}

	testCases := map[string]testData{
		"5つの要素がある場合": {
			inputs: [5]int{1, 2, 3, 4, 5},
			output: [5]int{5, 4, 3, 2, 1},
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			Reverse(&tc.inputs)

			if diff := cmp.Diff(tc.inputs, tc.output); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
