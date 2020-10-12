package ex5_1

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFindLinks(t *testing.T) {
	testDataDir, err := filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		fname         string
		wantErr       bool
		wantOut       []string
		wantErrString string
	}{
		"golang.orgの場合": {
			fname: filepath.Join(testDataDir, "golang_org.html"),
			wantOut: []string{
				"https://support.eji.org/give/153413/#!/donation/checkout",
				"/",
				"/doc/",
				"/pkg/",
				"/project/",
				"/help/",
				"/blog/",
				"https://play.golang.org/",
				"/dl/",
				"https://tour.golang.org/",
				"https://blog.golang.org/",
				"/doc/copyright.html",
				"/doc/tos.html",
				"http://www.google.com/intl/en/policies/privacy/",
				"http://golang.org/issues/new?title=x/website:",
				"https://google.com",
			},
		},
		"リンクが存在しない": {
			fname:   filepath.Join(testDataDir, "example.html"),
			wantOut: []string{},
		},
		"存在しないファイルが指定された時": {
			fname:         "notfound.html",
			wantErr:       true,
			wantErrString: "notfound.html is not found",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out, err := FindLInks(tc.fname)
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

			if diff := cmp.Diff(out, tc.wantOut); diff != "" {
				t.Errorf("mismatch output, diff: %v\n", diff)
			}
		})
	}
}
