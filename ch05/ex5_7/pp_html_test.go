package ex5_7

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"golang.org/x/net/html"
)

func TestPPrint(t *testing.T) {
	testDataDir, err := filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		fileName      string
		wantErr       bool
		wantErrString string
	}{
		"HTMLをパースできるかのテスト": {
			fileName: filepath.Join(testDataDir, "test.html"),
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
			pprinter := NewHTMLPrettyPrinter()
			b := &bytes.Buffer{}
			err = pprinter.PPrint(b, tc.fileName)
			if err != nil && !tc.wantErr {
				t.Errorf("this case not occur error, but occurred: %v\n", err)
			} else if tc.wantErr && err == nil {
				t.Errorf("error was expected, but error not occurred\n")
			} else if tc.wantErr && err != nil {
				if diff := cmp.Diff(err.Error(), tc.wantErrString); diff != "" {
					t.Errorf("mismatch error output, diff: %v\n", diff)
				}
			}

			_, err = html.Parse(bytes.NewReader(b.Bytes()))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
