package ex5_15

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMin(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		arguments       []int
		expectOut       int
		expectErr       bool
		expectErrString string
	}{
		"複数の引数が存在する": {
			arguments: []int{5, 4, 3, 2, 1},
			expectOut: 1,
		},
		"引数が存在しない": {
			arguments:       []int{},
			expectOut:       0,
			expectErr:       true,
			expectErrString: "arguments length is 0",
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out, err := Min(tc.arguments...)
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

func TestMax(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		arguments       []int
		expectOut       int
		expectErr       bool
		expectErrString string
	}{
		"複数の引数が存在する": {
			arguments: []int{1, 2, 3, 4, 5},
			expectOut: 5,
		},
		"引数が存在しない": {
			arguments:       []int{},
			expectOut:       0,
			expectErr:       true,
			expectErrString: "arguments length is 0",
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out, err := Max(tc.arguments...)
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
