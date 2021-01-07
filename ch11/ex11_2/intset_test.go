package intset

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func prepareIntSets() map[string]IntSet {
	return map[string]IntSet{
		"wordIntSet": &WordIntSet{},
		"mapIntSet":  NewMapIntSet(),
	}
}

func TestIntSet(t *testing.T) {
	testCases := prepareIntSets()

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			for _, elem := range []int{1, 2, 3, 4, 5} {
				tc.Add(elem)
			}

			if diff := cmp.Diff(tc.Len(), 5); diff != "" {
				t.Log(testName)
				t.Fatalf("invalid inset length, diff: %v", diff)
			}

			tc.Remove(5)

			if diff := cmp.Diff(tc.String(), "{1 2 3 4}"); diff != "" {
				t.Fatalf("invalid words, diff = %v", diff)
			}

			newIntSet := tc.Copy()

			if diff := cmp.Diff(tc.String(), newIntSet.String()); diff != "" {
				t.Fatalf("intset Copy invalid output, diff = %v", diff)
			}

		})
	}
}
