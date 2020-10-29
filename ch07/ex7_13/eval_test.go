package eval

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		//!-Eval
		// additional tests that don't appear in the book
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		//!+Eval
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestErrors(t *testing.T) {
	for _, test := range []struct{ expr, wantErr string }{
		{`"waiwai"`, "unexpected '\"'"},
		{"math.e", "unexpected '.'"},
		{"!true", "unexpected '!'"},
		{"x % 2", "unexpected '%'"},
		{"log(64)", `unknown function "log"`},
		{"pow(1)", "call to pow has 1 args, want 2"},
	} {
		expr, err := Parse(test.expr)
		if err == nil {
			vars := make(map[Var]bool)
			err = expr.Check(vars)
			if err == nil {
				t.Errorf("unexpected success: %s", test.expr)
				continue
			}
		}
		fmt.Printf("%-20s%v\n", test.expr, err) // (for book)
		if err.Error() != test.wantErr {
			t.Errorf("got error %s, want %s", err, test.wantErr)
		}
	}
}

func TestEvaluate(t *testing.T) {
	testCases := map[string]struct {
		expr   string
		env    Env
		expect string
	}{
		"足し算": {
			"x + y",
			Env{"x": 1, "y": 1},
			"2",
		},
		"引き算": {
			"a - b",
			Env{"a": 10, "b": 20},
			"-10",
		},
		"掛け算": {
			"i * j",
			Env{"i": 2, "j": 10},
			"20",
		},
		"割り算": {
			"m / n",
			Env{"m": 3, "n": 2},
			"1.5",
		},
		"根号": {
			"sqrt(A)",
			Env{"A": 64},
			"8",
		},
		"累乗": {
			"pow(n, m)",
			Env{"n": 2, "m": 10},
			"1024",
		},
		"複数の計算 1": {
			"A * B + (C - D)",
			Env{"A": 1, "B": 5, "C": 8, "D": 3},
			"10",
		},
		"複数の計算 2": {
			"sqrt(A) + pow(x, y)",
			Env{"A": 81, "x": 1, "y": 20},
			"10",
		},
	}

	var prev string
	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			if tc.expr != prev {
				t.Logf("\n\"%v\"\n", tc.expr)
				prev = tc.expr
			}

			e, err := Parse(tc.expr)
			if err != nil {
				t.Error(err)
			}

			got := fmt.Sprintf("%.5v", e.Eval(tc.env))
			t.Logf("%v --> %v\n", tc.env, got)

			if diff := cmp.Diff(got, tc.expect); diff != "" {
				t.Errorf("%v.Eval() in %v = %v, unexpected diff = %v", tc.expr, tc.env, got, diff)
			}
		})
	}
}
