package ex7_4

import (
	"io/ioutil"
	"testing"

	"golang.org/x/net/html"

	"github.com/google/go-cmp/cmp"
)

func TestStringReader(t *testing.T) {
	t.Parallel()

	sr := NewReader("test string")

	teststr1 := "test1"

	read, err := sr.Read([]byte(teststr1))
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(read, len([]byte(teststr1))); diff != "" {
		t.Fatalf("invalid output length diff = %v", diff)
	}
}

func TestStringReaderHTML(t *testing.T) {
	t.Parallel()

	b, err := ioutil.ReadFile("./htmlFile/index.html")
	if err != nil {
		t.Fatal(err)
	}

	_, err = html.Parse(NewReader(string(b)))
	if err != nil {
		t.Fatal(err)
	}
}
