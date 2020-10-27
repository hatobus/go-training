package countingwriter

import "io"

type ByteCounter struct {
	w       io.Writer
	written int64
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	n, err := bc.w.Write(p)
	if err != nil {
		return 0, err
	}

	bc.written += int64(n)

	return n, nil
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &ByteCounter{w, 0}
	return c, &c.written
}
