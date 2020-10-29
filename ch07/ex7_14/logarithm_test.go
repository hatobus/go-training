package eval

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLog(t *testing.T) {
	type exp struct {
		expr string
		env  Env
	}

	for _, tc := range []struct {
		testName string
		exp      exp
		expect   float64
	}{
		{
			testName: "log2(8)",
			exp: exp{
				"log2(8)",
				Env{},
			},
			expect: 3,
		},
		{
			testName: "log10(1000)",
			exp: exp{
				"log10(1000)",
				Env{},
			},
			expect: 3,
		},
		{
			testName: "log2(x)",
			exp: exp{
				expr: "log2(x)",
				env:  Env{"x": 8},
			},
			expect: 3,
		},
		{
			testName: "log10(x)",
			exp: exp{
				expr: "log10(y)",
				env:  Env{"y": 1000},
			},
			expect: 3,
		},
		{
			testName: "-log2(x)",
			exp: exp{
				expr: "-log2(x)",
				env:  Env{"x": 8},
			},
			expect: -3,
		},
		{
			testName: "-log10(x)",
			exp: exp{
				expr: "-log10(x)",
				env:  Env{"x": 1000},
			},
			expect: -3,
		},
		{
			testName: "log2(x)-log10(y)",
			exp: exp{
				expr: "log2(x)-log10(y)",
				env:  Env{"x": 8, "y": 1000},
			},
			expect: 0,
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			expr, err := Parse(tc.exp.expr)
			if err != nil {
				t.Error(err)
			}

			out := expr.Eval(tc.exp.env)

			if diff := cmp.Diff(out, tc.expect); diff != "" {
				t.Errorf("invalid output diff = %v", diff)
			}
		})
	}
}
