package e2e_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	main "github.com/hatobus/go-training/ch04/ex4_5/ex4.14/cmd"
	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/github"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/e2e/fakeserver"
)

var fakeIssueListServerURL string
var mainArgs = []string{"hatobus", "go-training"}

func TestMain(m *testing.M) {
	wantIssuelistPathParam := strings.Join([]string{"/repos", mainArgs[0], mainArgs[1], "issues"}, "/")
	issueListServer := httptest.NewServer(fakeserver.FakeIssueListAPIHandler(wantIssuelistPathParam))
	defer issueListServer.Close()

	fakeIssueListServerURL = issueListServer.URL

	m.Run()
}

func TestMainServerIssueListHandler(t *testing.T) {
	github.APIURL = fakeIssueListServerURL

	// ここで 8080 が開く
	main.RunMain(t, mainArgs)

	// localhost:8080/list にアクセスしてHTMLを取得
	resp, err := http.Get("http://localhost:8080/list")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	htmlbyte, err := ioutil.ReadFile("./testData/HTML/issue_list.html")
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(string(b), string(htmlbyte)); diff != "" {
		t.Fatalf("invalid html, diff = %v", diff)
	}
}
