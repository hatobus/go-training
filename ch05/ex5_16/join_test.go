package ex5_16

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestJoin(t *testing.T) {
	testData := map[string]struct {
		arguments       []string
		sep             string
		expectOut       string
		expectErr       bool
		expectErrString string
	}{
		"複数の引数が存在する": {
			arguments: []string{"a", "b", "c"},
			sep:       ", ",
			expectOut: "a, b, c",
		},
		"引数が1つの場合": {
			arguments: []string{"a"},
			sep:       ", ",
			expectOut: "a",
		},
		"引数が無い場合": {
			arguments:       []string{},
			sep:             ", ",
			expectOut:       "",
			expectErr:       true,
			expectErrString: "argument length is 0",
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out, err := Join(tc.sep, tc.arguments...)
			if err != nil {
				switch tc.expectErr {
				case true:
					if diff := cmp.Diff(err.Error(), tc.expectErrString); diff != "" {
						t.Fatalf("invalid error string diff: %v", diff)
					}
				default:
					t.Fatalf("error unexpected, but error occured error: %v", err)
				}
			}

			if diff := cmp.Diff(out, tc.expectOut); diff != "" {
				t.Fatalf("unexpected output: %v", diff)
			}
		})
	}
}
