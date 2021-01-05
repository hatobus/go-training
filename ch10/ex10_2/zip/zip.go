package zip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"

	"github.com/hatobus/go-training/ch10/ex10_2/file_printer"
)

type reader struct {
	Reader *zip.Reader
	files  []*zip.File
	r      io.ReadCloser
	write  string
}

func init() {
	file_printer.RegisterFormat("zip", "PK", NewReader)
}

func (r *reader) Read(b []byte) (int, error) {
	if r.r == nil && len(r.files) == 0 {
		return 0, io.EOF
	}

	if r.r == nil {
		// 次のファイルを読み込む
		f := r.files[0]
		r.files = r.files[1:]
		var err error
		r.r, err = f.Open()
		if err != nil {
			return 0, fmt.Errorf("failed to read zip file: %v", err)
		}
		if f.Mode()&os.ModeDir == 0 {
			r.write = f.Name + ":\n"
		}
	}

	writtenSize := 0

	if len(r.write) > 0 {
		n := copy(b, r.write)
		b = b[n:]
		r.write = r.write[n:]
		writtenSize += n
	}

	n, err := r.r.Read(b)
	writtenSize += n
	if err != nil {
		r.r.Close()
		r.r = nil
		if err == io.EOF {
			return writtenSize, nil
		}
	}

	return writtenSize, nil
}

func NewReader(f *os.File) (io.Reader, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("zip reader failed: %v", err)
	}
	r, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return nil, fmt.Errorf("zip reader failed: %v", err)
	}
	return &reader{r, r.File, nil, ""}, nil
}
