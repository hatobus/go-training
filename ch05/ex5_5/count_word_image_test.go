package ex5_5

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testDataDir string

func getHTMLHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "invalid request method", http.StatusBadRequest)
			return
		}

		fname := req.URL.Query().Get("filename")
		if fname == "" {
			http.Error(w, "filename expected but not found", http.StatusBadRequest)
			return
		}

		fname = filepath.Join(testDataDir, fname)

		_, err := os.Stat(fname)
		if os.IsNotExist(err) {
			http.Error(w, "your file not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Open(fname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func TestCountWordsAndImage(t *testing.T) {
	var err error
	testDataDir, err = filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		fileName      string
		wantErr       bool
		expectWords   int
		expectImages  int
		wantErrString string
	}{
		"画像の存在しないHTMLファイルの場合": {
			fileName:     "href.html",
			expectWords:  22,
			expectImages: 0,
		},
		"画像と文字のHTML": {
			fileName:     "multi_links.html",
			expectWords:  11,
			expectImages: 2,
		},
		"存在しないファイル": {
			fileName:      "not_found.html",
			wantErr:       true,
			wantErrString: "your requested URL returned status code 2XX, returned 404",
		},
	}

	ts := httptest.NewServer(getHTMLHandler())
	t.Cleanup(func() {
		ts.Close()
	})

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			u, err := url.Parse(ts.URL)
			if err != nil {
				t.Fatal(err)
			}

			q := u.Query()
			q.Set("filename", tc.fileName)
			u.RawQuery = q.Encode()

			words, images, err := CountWordsAndImage(u.String())
			if err != nil && !tc.wantErr {
				t.Errorf("this case not occur error, but occurred: %v\n", err)
			} else if tc.wantErr && err == nil {
				t.Errorf("error was expected, but error not occurred\n")
			} else if tc.wantErr && err != nil {
				if diff := cmp.Diff(err.Error(), tc.wantErrString); diff != "" {
					t.Errorf("mismatch error output, diff: %v\n", diff)
				}
			}

			if diff := cmp.Diff(words, tc.expectWords); diff != "" {
				t.Errorf("mismatch expect words, diff: %v\n", diff)
			}

			if diff := cmp.Diff(images, tc.expectImages); diff != "" {
				t.Errorf("mismatch expect images, diff: %v\n", diff)
			}
		})
	}
}
