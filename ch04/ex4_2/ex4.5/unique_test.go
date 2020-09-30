package unique

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniqueStrings(t *testing.T) {
	type testData struct {
		inp []string
		out []string
	}

	testCases := map[string]testData{
		"1つ重複がある": {
			inp: []string{"a", "a", "b", "c"},
			out: []string{"a", "b", "c"},
		},
		"重複がない": {
			inp: []string{"a", "b", "c"},
			out: []string{"a", "b", "c"},
		},
		"複数の重複": {
			inp: []string{"a", "a", "b", "a", "c", "c"},
			out: []string{"a", "b", "a", "c"},
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			unique := UniqueStrings(tc.inp)

			if diff := cmp.Diff(tc.out, unique); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
