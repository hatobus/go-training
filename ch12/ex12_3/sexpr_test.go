package ex12_3

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalling(t *testing.T) {
	testCases := map[string]struct {
		input  interface{}
		expect string
	}{
		"complex value": {
			input:  0 + 1i,
			expect: "#C(0 1)",
		},
		"boolean true": {
			input:  true,
			expect: "t",
		},
		"boolean false": {
			input:  false,
			expect: "nil",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			marshaled, err := Marshal(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expect, string(marshaled)); diff != "" {
				t.Fatalf("invalid output, diff: %v\n", diff)
			}
		})
	}
}
