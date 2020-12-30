package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
)

var (
	format   = flag.String("format", "png", "format with output ")
	filename = flag.String("name", "out", "output filename")
)

func main() {
	flag.Parse()

	if !(*format == "png" || *format == "gif" || *format == "jpg") {
		fmt.Fprintf(os.Stderr, "invalid format, format need to \"png\" or \"jpg\" or \"gif\"")
		os.Exit(1)
	}

	var saveFileName string
	ext := filepath.Ext(*filename)
	if ext != *format {
		saveFileName = filepath.Dir(*filename) + filepath.Base(*filename) + "." + ext
	} else {
		saveFileName = *filename
	}

	wd, _ := os.Getwd()
	savedir := filepath.Join(wd, saveFileName)

	f, err := os.Open(savedir)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	if err := toJPEG(os.Stdin, f); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
