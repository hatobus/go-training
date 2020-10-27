package ex7_4

import "io"

// Reader is the interface that wraps the basic Read method
// https://golang.org/pkg/io/#Reader

//type Reader interface {
//	Read(p []byte) (n int, err error)
//}

type StringReader struct {
	str string
}

func (sr *StringReader) Read(p []byte) (int, error) {
	n := copy(p, sr.str)
	sr.str = sr.str[n:]

	if len(sr.str) == 0 {
		return 0, io.EOF
	}

	return n, nil
}

func NewReader(s string) io.Reader {
	return &StringReader{s}
}
