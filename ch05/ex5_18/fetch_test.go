package ex5_18

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var origHTMLFilesDir string
var saveDataDir = "saveData"

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func getHTMLHandler(t testing.TB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			t.Log("invalid request method")
			http.Error(w, "invalid request method", http.StatusBadRequest)
			return
		}

		t.Log(req.URL.Path)

		expectFilePath, _ := shiftPath(req.URL.Path)
		fname := path.Base(expectFilePath)
		if fname == "" {
			t.Log("expect filename not found")
			http.Error(w, "filename expected but not found", http.StatusBadRequest)
			return
		}

		fname = filepath.Join(origHTMLFilesDir, fname)

		_, err := os.Stat(fname)
		if os.IsNotExist(err) {
			t.Logf("file %v not found in directory\n", fname)
			http.Error(w, "your file not found", http.StatusNotFound)
			return
		} else if err != nil {
			t.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		f, err := os.Open(fname)
		if err != nil {
			t.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Log(fname)
			t.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func TestFetch(t *testing.T) {
	var err error
	origHTMLFilesDir, err = filepath.Abs("./orig_htmls")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(getHTMLHandler(t))
	t.Cleanup(func() {
		ts.Close()
	})

	var saveDirPath string
	c, err := filepath.Abs("./")
	if err != nil {
		t.Fatal(err)
	}

	saveDirPath = filepath.Join(c, ts.URL, saveDataDir)

	testCases := map[string]struct {
		fileName            string
		expectFilePath      string
		expectNumberOfBytes int64
		expectErr           bool
		expectErrString     string
	}{
		"指定したファイルを取ってローカルに保存する": {
			fileName:            "one.html",
			expectFilePath:      filepath.Join(saveDirPath, "href.html"),
			expectNumberOfBytes: 183,
		},
		"index.htmlを取ってくる": {
			fileName:            "",
			expectFilePath:      filepath.Join(saveDataDir, "index.html"),
			expectNumberOfBytes: 332,
		},
		"存在しないファイルを取ってこようとする": {
			fileName:        "notFound.html",
			expectErr:       true,
			expectErrString: "resposne status code 404",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			u, err := url.Parse(ts.URL)
			if err != nil {
				t.Fatal(err)
			}

			if tc.fileName != "" {
				u.Path = path.Join(u.Path, tc.fileName)
			}

			fileName, length, err := Fetch(u.String(), saveDataDir)
			if err != nil && !tc.expectErr {
				t.Errorf("this case not occur error, but occurred: %v\n", err)
			} else if tc.expectErr && err == nil {
				t.Errorf("error was expected, but error not occurred\n")
			} else if tc.expectErr && err != nil {
				if diff := cmp.Diff(err.Error(), tc.expectErrString); diff != "" {
					t.Errorf("mismatch error output, diff: %v\n", diff)
				}
			}

			t.Log(fileName)

			if diff := cmp.Diff(length, tc.expectNumberOfBytes); diff != "" {
				t.Fatalf("invalid expect file length, diff = %v\n", diff)
			}

			var fn string
			fn = tc.fileName
			if fn == "" {
				fn = "index.html"
			}
			origDataFullPath := filepath.Join(origHTMLFilesDir, fn)

			downloadFileFullPath, err := filepath.Abs(filepath.Join(".", fileName))
			if err != nil {
				t.Fatal(err)
			}

			origData, err := ioutil.ReadFile(origDataFullPath)
			if err != nil {
				t.Fatal(err)
			}

			downloadFileData, err := ioutil.ReadFile(downloadFileFullPath)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(string(origData), string(downloadFileData)); diff != "" {
				t.Fatalf("%v and %v is not same, diff: %v", origDataFullPath, tc.expectFilePath, diff)
			}
		})
	}
}
