package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var origHTMLFilesDir string

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func generateIndexHTML(t testing.TB, rawurl string) {
	tpl := template.Must(template.ParseFiles("util/index.html"))

	f, err := os.Create(filepath.Join(origHTMLFilesDir, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = tpl.Execute(f, rawurl)
	if err != nil {
		t.Fatal(err)
	}
}

func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func getLastPathParameter(s string) string {
	p := strings.Split(s, "/")
	return p[len(p)-1]
}

func getHTMLHandler(t testing.TB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			t.Log("invalid request method")
			http.Error(w, "invalid request method", http.StatusBadRequest)
			return
		}

		expectFilePath, _ := shiftPath(req.URL.Path)
		fname := getLastPathParameter(expectFilePath)
		if fname == "" {
			t.Log("expect filename not found")
			http.Error(w, "filename expected but not found", http.StatusBadRequest)
			return
		}

		fname = filepath.Join(origHTMLFilesDir, fname)

		_, err := os.Stat(fname)
		if os.IsNotExist(err) {
			t.Log("file not found in directory")
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
			t.Log(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func getLastPath(rawurl string) string {
	component := strings.Split(rawurl, "/")
	return component[len(component)-1]
}

func TestCrawl(t *testing.T) {
	var err error
	origHTMLFilesDir, err = filepath.Abs("./orig_htmls")
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(getHTMLHandler(t))
	defer ts.Close()

	hostName := getLastPath(ts.URL)

	replicaDir, err := filepath.Abs(filepath.Join("./replica_html", hostName))
	if err != nil {
		t.Fatal(err)
	}

	defer os.RemoveAll(replicaDir)

	generateIndexHTML(t, ts.URL)

	t.Log(ts.URL)

	err = BreadthFirst(nil, []string{ts.URL + "/index.html"})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(replicaDir)

	replicaFiles := dirwalk(replicaDir)
	replicaMap := make(map[string]string, 0)

	for _, fullPath := range replicaFiles {
		replicaMap[getLastPath(fullPath)] = fullPath
	}

	origFiles := dirwalk(origHTMLFilesDir)
	origMap := make(map[string]string, 0)

	for _, fullPath := range origFiles {
		origMap[getLastPath(fullPath)] = fullPath
	}

	for fname, fullPath := range replicaMap {
		origFullPath, ok := origMap[fname]
		if !ok {
			t.Fatalf("file %v not found in original file", fname)
		}

		origData, err := ioutil.ReadFile(origFullPath)
		if err != nil {
			t.Fatal(err)
		}

		replicaData, err := ioutil.ReadFile(fullPath)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(string(origData), string(replicaData)); diff != "" {
			t.Fatalf("%v and %v is not same, diff: %v", origFullPath, fullPath, diff)
		}
	}
}
