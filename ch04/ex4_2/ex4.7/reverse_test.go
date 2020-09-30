package reverse

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverse(t *testing.T) {
	type testData struct {
		input []byte
		want  string
	}

	testCases := map[string]testData{
		"英数字": {
			input: []byte("test"),
			want:  "tset",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			ByteReverse([]byte(tc.input))

			if diff := cmp.Diff(string(tc.input), tc.want); diff != "" {
				t.Fatalf("invalid output , diff: %v", diff)
			}
		})
	}
}
