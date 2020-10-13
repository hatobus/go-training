package ex5_2

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHTMLTagFrequency(t *testing.T) {
	testDataDir, err := filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		fileName      string
		wantErr       bool
		expectOut     map[string]int
		wantErrString string
	}{
		"golang.orgの場合": {
			fileName: filepath.Join(testDataDir, "golang_org.html"),
			expectOut: map[string]int{
				"a":        34,
				"body":     1,
				"br":       1,
				"button":   8,
				"div":      28,
				"footer":   2,
				"form":     2,
				"h1":       2,
				"h2":       6,
				"header":   2,
				"html":     1,
				"i":        2,
				"iframe":   2,
				"img":      3,
				"input":    1,
				"li":       22,
				"link":     4,
				"main":     2,
				"meta":     4,
				"nav":      2,
				"noscript": 2,
				"option":   16,
				"p":        2,
				"path":     2,
				"pre":      2,
				"script":   16,
				"section":  8,
				"select":   2,
				"strong":   6,
				"svg":      2,
				"textarea": 2,
				"title":    4,
				"ul":       4,
			},
		},
		"h1からh6までの要素を持ったHTMLファイル": {
			fileName: filepath.Join(testDataDir, "headings.html"),
			expectOut: map[string]int{
				"html": 2,
				"body": 2,
				"h1":   2,
				"h2":   2,
				"h3":   2,
				"h4":   2,
				"h5":   2,
				"h6":   2,
			},
		},
		"存在しないファイルが指定された時": {
			fileName:      "notfound.html",
			wantErr:       true,
			wantErrString: "notfound.html is not found",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out, err := HTMLTagFrequency(tc.fileName)
			if err != nil && !tc.wantErr {
				t.Errorf("this case not occur error, but occurred: %v\n", err)
			} else if tc.wantErr && err == nil {
				t.Errorf("error was expected, but error not occurred\n")
			} else if tc.wantErr && err != nil {
				if diff := cmp.Diff(err.Error(), tc.wantErrString); diff != "" {
					t.Errorf("mismatch error output, diff: %v\n", diff)
				}
			}

			if diff := cmp.Diff(out, tc.expectOut); diff != "" {
				t.Errorf("mismatch output, diff: %v\n", diff)
			}
		})
	}
}
