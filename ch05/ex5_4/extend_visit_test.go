package ex5_4

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestExtractLinksFromHTML(t *testing.T) {
	testDataDir, err := filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		fileName      string
		wantErr       bool
		expectOut     []string
		wantErrString string
	}{
		"href要素のURLのみの場合": {
			fileName: filepath.Join(testDataDir, "href.html"),
			expectOut: []string{
				"https://www.w3schools.com",
			},
		},
		"img要素を持ったHTML": {
			fileName: filepath.Join(testDataDir, "img.html"),
			expectOut: []string{
				"html5.gif",
				"internet explorer.gif",
			},
		},
		"複数のリンクが存在するHTML": {
			fileName: filepath.Join(testDataDir, "multi_links.html"),
			expectOut: []string{
				"html5.gif",
				"internet explorer.gif",
				"sample.js",
				"stylesheet",
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

			out, err := ExtractLinksFromHTML(tc.fileName)
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
