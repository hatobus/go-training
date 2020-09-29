package anagram

import "testing"

func TestAnagram(t *testing.T) {
	type testData struct {
		s1   string
		s2   string
		same bool
	}

	testCases := map[string]testData{
		"英語でアナグラムを検知できる": {
			s1:   "abc",
			s2:   "cba",
			same: true,
		},
		"英語でアナグラムでない時": {
			s1:   "abc",
			s2:   "aab",
			same: false,
		},
		"長い英文(完全パングラム)": {
			s1:   "Jumbling vext frowzy hacks PDQ",
			s2:   "Glum Schwartzkopf vexd by NJ IQ",
			same: true,
		},
		"日本語でアナグラム": {
			s1:   "こんにちは",
			s2:   "こんちには",
			same: true,
		},
		"日本語でアナグラムでない時": {
			s1:   "こんにちは",
			s2:   "こんばんは",
			same: false,
		},
		"日本語でアナグラム(サロゲートペア)": {
			s1:   "𠮷野家で𩸽",
			s2:   "𩸽で𠮷野家",
			same: true,
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			isSame := Anagram(tc.s1, tc.s2)

			if isSame != tc.same {
				t.Logf("invalid testcase, s1: %v, s2: %v", tc.s1, tc.s2)
				t.Fatalf("invalid output, want %v but got %v", tc.same, isSame)
			}
		})
	}
}
