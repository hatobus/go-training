package tar

import (
	"archive/tar"
	"io"
	"os"

	"github.com/hatobus/go-training/ch10/ex10_2/file_printer"
)

type reader struct {
	Reader *tar.Reader
	file   *os.File
	write  string
}

func init() {
	file_printer.RegisterFormat("tar", "ustar", NewReader)
}

func (r *reader) Read(b []byte) (int, error) {
	writtenSize := 0
	for len(b) > 0 {
		if len(r.write) > 0 {
			n := copy(b, r.write)
			writtenSize += n
			r.write = r.write[n:]
			b = b[n:]
		}

		n, err := r.Reader.Read(b)
		writtenSize += n
		b = b[n:]

		switch err {
		case nil:
			continue
		case io.EOF:
			// 次のファイルを見る
			next, err := r.Reader.Next()
			if err != nil {
				return writtenSize, err
			}
			if next.Typeflag == tar.TypeDir {
				continue
			}
			r.write = next.Name + ":\n"
		default:
			return writtenSize, err
		}
	}
	return writtenSize, nil
}

func NewReader(f *os.File) (io.Reader, error) {
	return &reader{tar.NewReader(f), f, ""}, nil
}
