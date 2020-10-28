package kaibun

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestKaibun(t *testing.T) {
	testCases := map[string]struct {
		values []int
		expect bool
	}{
		"回分になっている(要素が偶数個)": {
			values: []int{1, 1, 2, 2, 1, 1},
			expect: true,
		},
		"回分になっている(要素が奇数個)": {
			values: []int{1, 1, 2, 1, 1},
			expect: true,
		},
		"回分になっていない": {
			values: []int{1, 1, 4, 5, 1, 4},
			expect: false,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(tc.expect, IsKaibun(sort.IntSlice(tc.values))); diff != "" {
				t.Fatalf("invalid output diff = %v", diff)
			}
		})
	}
}
