package ex12_13

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMarshalling(t *testing.T) {
	type Struct struct {
		Bool       bool
		Complex64  complex64 `sexpr:"complex"`
		Complex128 complex128
		Float32    float32
		Float64    float64 `sexpr:"float"`
		Interface  interface{}
	}
	testCases := map[string]struct {
		input  Struct
		expect string
	}{
		"すべての値が入っている": {
			input: Struct{
				Bool:       true,
				Complex64:  complex(0, 1),
				Complex128: complex(2, 3),
				Float32:    4.5,
				Float64:    6.78,
				Interface:  "interface",
			},
			expect: `((Bool t) (complex #C(0 1)) (Complex128 #C(2 3)) (Float32 4.5) (float 6.78) (Interface ("interface {}" "interface")))`,
		},
		"一部のフィールドが空": {
			input: Struct{
				Complex64: complex(0, 1),
				Float64:   2.345,
			},
			expect: `((Bool nil) (complex #C(0 1)) (Complex128 #C(0 0)) (Float32 0) (float 2.345) (Interface ("interface {}" nil)))`,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			d, err := Marshal(tc.input)
			t.Log(string(d))
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(tc.expect, string(d)); diff != "" {
				t.Errorf("unexpected output, diff = %v", diff)
			}
		})
	}
}

