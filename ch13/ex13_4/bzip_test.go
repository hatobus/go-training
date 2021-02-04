package ex13_4

import (
	"bytes"
	"compress/bzip2"
	"io"
	"testing"
)

func TestBzip(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w, err := NewWriter(&compressed)
	if err != nil {
		t.Error(err)
		return
	}

	mw := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 10000; i++ {
		io.WriteString(mw, "aaaaa")
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}

	var uc bytes.Buffer
	io.Copy(&uc, bzip2.NewReader(&compressed))
	if !bytes.Equal(uncompressed.Bytes(), uc.Bytes()) {
		t.Error("failed to decompression")
	}
}
