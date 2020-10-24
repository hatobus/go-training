package main

import (
	"fmt"
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

func TestOutline(t *testing.T) {
	var err error

	testDataDir, err = filepath.Abs("./testData")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		filename        string
		expectOut       string
		expectErr       bool
		expectErrString string
	}{
		"通常のHTMLファイル": {
			filename:  "plain.html",
			expectOut: fmt.Sprintf("</html>\n  </head>\n  </head>\n  </body>\n    </h2>\n    </h2>\n    </p>\n    </p>\n    </a>\n    </a>\n  </body>\n</html>\n"),
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
			q.Set("filename", tc.filename)
			u.RawQuery = q.Encode()

			out, err := outline(u.String())
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expectOut, out); diff != "" {
				t.Fatalf("invalid output, diff: %v", diff)
			}
		})
	}
}
