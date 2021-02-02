package ex12_2

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDisplay(t *testing.T) {
	type Pointer *Pointer
	var p Pointer
	p = &p

	type S struct {
		name string
		val  *S
	}
	var s S
	s = S{"cycle", &s}

	testCases := map[string]struct {
		input  interface{}
		expect string
	}{
		"ポインタのポインタ": {
			input:  p,
			expect: fmt.Sprintf("(*(*(*(*input)))) = ex12_2.Pointer %p\n", p),
		},
		"structの要素に自分自身が入る": {
			input:  s,
			expect: "input.name = \"cycle\"\n(*input.val).name = \"cycle\"\n(*(*input.val).val) = ex12_2.S value\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			displayed := Display("input", tc.input)
			if diff := cmp.Diff(tc.expect, displayed); diff != "" {
				t.Errorf("invalid output: diff = %v\n", diff)
			}
		})
	}
}
