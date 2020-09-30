package u2a

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUnicode2AsciiTrimSpace(t *testing.T) {
	type testData struct {
		input string
		want  string
	}

	testCases := map[string]testData{
		"複数のスペースが入っている時": testData{
			input: `Hoge  Fuga     Foo`,
			want:  `Hoge Fuga Foo`,
		},
		"複数のスペースが入っている時(制御文字)": testData{
			input: fmt.Sprint("Hoge\tFuga\t\tFoo"),
			want:  `Hoge Fuga Foo`,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := Unicode2AsciiTrimSpace([]byte(tc.input))

			if diff := cmp.Diff(string(out), tc.want); diff != "" {
				t.Fatalf("invalid output, diff %v", diff)
			}
		})
	}
}
