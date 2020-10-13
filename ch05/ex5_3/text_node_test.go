package ex5_3

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHTMLTextNodes(t *testing.T) {
	testDataDir, err := filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		fileName      string
		wantErr       bool
		expectOut     map[string][]string
		wantErrString string
	}{
		"h1からh6までの要素を持ったHTMLファイル": {
			fileName: filepath.Join(testDataDir, "headings.html"),
			expectOut: map[string][]string{
				"h1": {"This is heading 1"},
				"h2": {"This is heading 2"},
				"h3": {"This is heading 3"},
				"h4": {"This is heading 4"},
				"h5": {"This is heading 5"},
				"h6": {"This is heading 6"},
			},
		},
		"scriptタグとstyleタグは無視する": {
			fileName: filepath.Join(testDataDir, "ignore_tag.html"),
			expectOut: map[string][]string{
				"noscript": {
					"Sorry, your browser does not support JavaScript!",
					"noscript tag can't ignore",
				},
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

			out, err := HTMLTextNodes(tc.fileName)
			t.Log(out)
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
