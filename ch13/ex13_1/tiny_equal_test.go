package ex13_1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTinyEqual(t *testing.T) {
	testCases := map[string]struct {
		x      interface{}
		y      interface{}
		expect bool
	}{
		"Int & Int => True": {
			x:      1,
			y:      1,
			expect: true,
		},
		"Int & Int => False": {
			x:      1,
			y:      0,
			expect: false,
		},
		"Int & Float => False": {
			x:      1,
			y:      0.1,
			expect: false,
		},
		"Float & Float => True": {
			x:      1000000000.1234,
			y:      1000000000.0000,
			expect: true,
		},
		"Float & Float => False": {
			x:      100000000,
			y:      100000001,
			expect: false,
		},
		"String & String => True": {
			x:      "true",
			y:      "true",
			expect: true,
		},
		"String & String => False": {
			x:      "true",
			y:      "false",
			expect: false,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			tf := Equal(tc.x, tc.y)
			if diff := cmp.Diff(tf, tc.expect); diff != "" {
				t.Fatalf("unexpected output, diff = %v", diff)
			}
		})
	}
}
