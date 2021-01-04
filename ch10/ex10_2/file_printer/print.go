package file_printer

import (
	"fmt"
	"os"

	"github.com/hatobus/go-training/ch10/ex10_2/detector"
)

func FilePrint(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	buf := make([]byte, 32)
	_, err = f.Read(buf)
	if err != nil {
		return err
	}

	mime, err := detector.DetectFileType(buf)
	if err != nil {
		return err
	}

	fmt.Printf("your file is: %v\n", mime)

	return nil
}
