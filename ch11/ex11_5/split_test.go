package split

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplit(t *testing.T) {
	testCases := map[string]struct {
		s            string
		sep          string
		expectElems  []string
		expectLength int
	}{
		"デフォルトのテストケース": {
			"a:b:c",
			":",
			[]string{"a", "b", "c"},
			3,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			splited := strings.Split(tc.s, tc.sep)
			if diff := cmp.Diff(len(splited), tc.expectLength); diff != "" {
				t.Fatalf("invalid expect length, diff = %v", diff)
			}

			if diff := cmp.Diff(splited, tc.expectElems); diff != "" {
				t.Fatalf("invalid expect elems, diff = %v", diff)
			}
		})
	}
}
