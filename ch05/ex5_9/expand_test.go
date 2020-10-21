package ex5_9

import (
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExpand(t *testing.T) {
	testData := map[string]struct {
		srcString        string
		genReplaceString func(string) string
		expectString     string
	}{
		"引っかかった文字を大文字に変換する": {
			srcString: "abc$def",
			genReplaceString: func(s string) string {
				return strings.ToUpper(s)
			},
			expectString: "abcDEF",
		},
		"引っかかった文字を辞書順に並び替え": {
			srcString: "abc$dbaetc",
			genReplaceString: func(s string) string {
				ss := strings.Split(s, "")
				sort.Strings(ss)
				return strings.Join(ss, "")
			},
			expectString: "abcabcdet",
		},
		"引っかかった文字を削除する": {
			srcString: "abc$def",
			genReplaceString: func(s string) string {
				return ""
			},
			expectString: "abc",
		},
		"引っかからなかったらそのまま": {
			srcString: "abcdef",
			genReplaceString: func(s string) string {
				return strings.ToUpper(s)
			},
			expectString: "abcdef",
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := Expand(tc.srcString, tc.genReplaceString)

			if diff := cmp.Diff(tc.expectString, out); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
