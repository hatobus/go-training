package file_printer

import (
	"fmt"
	"io"
	"os"

	"github.com/hatobus/go-training/ch10/ex10_2/detector"
)

type NewReader func(*os.File) (io.Reader, error)

type Format struct {
	name        string
	magicNumber string
	reader      NewReader
}

var formats []Format

func RegisterFormat(name, magicNumber string, f NewReader) {
	formats = append(formats, Format{name, magicNumber, f})
}

func FilePrint(filename string) (io.Reader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 32)
	_, err = f.Read(buf)
	if err != nil {
		return nil, err
	}

	mime, err := detector.DetectFileType(buf)
	if err != nil {
		return nil, err
	}

	format := new(Format)
	for _, f := range formats {
		if f.name == mime {
			format = &f
			break
		}
	}

	fmt.Printf("your file is: %v\n", format)

	return format.reader(f)
}
