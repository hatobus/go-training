package sha256diff

import "testing"

func TestGetSHA256Diff(t *testing.T) {
	type testData struct {
		s1      string
		s2      string
		diffLen int
	}

	testCases := map[string]testData{
		"xとXの比較": {
			s1:      "x",
			s2:      "X",
			diffLen: 31,
		},
		"xとxの比較": {
			s1:      "x",
			s2:      "x",
			diffLen: 0,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := GetSHA256Diff(tc.s1, tc.s2)

			if out != tc.diffLen {
				t.Fatalf("invalid output, want %v, but got %v", tc.diffLen, out)
			}
		})
	}
}
