package countingwriter

import (
	"bytes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}

	c, n := CountingWriter(b)
	data := []byte("hi there")

	if _, err := c.Write(data); err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(*n, int64(len(data))); diff != "" {
		t.Fatalf("invalid length diff = %v", diff)
	}
}
