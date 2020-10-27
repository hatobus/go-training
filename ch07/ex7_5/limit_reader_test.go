package ex7_5

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLimitReaderLTLimit(t *testing.T) {
	t.Parallel()
	testStr := "test string"

	b := &bytes.Buffer{}

	lr := NewLimitReader(strings.NewReader(testStr), 20)

	n, err := b.ReadFrom(lr)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(int(n), len(testStr)); diff != "" {
		t.Fatalf("invalid read length diff = %v", diff)
	}

	if diff := cmp.Diff(b.String(), testStr); diff != "" {
		t.Fatalf("invalid string diff = %v", diff)
	}
}

func TestLimitReaderGTLimit(t *testing.T) {
	t.Parallel()
	testStr := "test string"

	b := &bytes.Buffer{}

	limit := 5

	lr := NewLimitReader(strings.NewReader(testStr), int64(limit))

	n, err := b.ReadFrom(lr)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(int(n), limit); diff != "" {
		t.Fatalf("invalid output diff = %v", diff)
	}

	if diff := cmp.Diff(b.String(), "test "); diff != "" {
		t.Fatalf("invalid string diff = %v", diff)
	}

	_, err = lr.Read([]byte("eof"))
	if err != io.EOF {
		t.Fatalf("unexpected error occured, want io.EOF but error %v", err)
	}
}
