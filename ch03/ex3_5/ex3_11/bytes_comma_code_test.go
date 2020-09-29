package comma

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBytesComma(t *testing.T) {
	type testData struct {
		input     string
		output    string
		wanterror bool
		errString string
	}

	testCases := map[string]testData{
		"3文字以下の時": {
			input:  "-10",
			output: "-10",
		},
		"コンマが1つ入る": {
			input:  "+1000",
			output: "+1,000",
		},
		"コンマが2つ入る": {
			input:  "1000000",
			output: "1,000,000",
		},
		"小数点がある場合": {
			input:  "0.05",
			output: "0.05",
		},
		"小数点がありコンマも存在する場合": {
			input:  "1234.56",
			output: "1,234.56",
		},
		"複数の小数点が存在する": {
			input:     "1.234.56",
			wanterror: true,
			errString: "invalid format",
		},
		".始まり": {
			input:     ".56",
			wanterror: true,
			errString: "invalid format",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			out, err := BytesCommaWithCode(tc.input)
			if err != nil {
				if !tc.wanterror {
					t.Fatalf("error occured, err: %v", err)
				} else if err.Error() != tc.errString {
					t.Fatalf("other error occured, want %v but got %v", tc.errString, err.Error())
				}
			}

			if diff := cmp.Diff(tc.output, out); diff != "" {
				t.Fatalf("invalid output, diff %v", diff)
			}
		})
	}
}
