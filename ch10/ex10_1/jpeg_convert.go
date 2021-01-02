package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
)

var (
	format   = flag.String("format", "png", "format with output ")
	filename = flag.String("name", "out.png", "output filename (rel path)")
)

func main() {
	flag.Parse()

	if !(*format == "png" || *format == "gif" || *format == "jpg") {
		fmt.Fprintf(os.Stderr, "invalid format, format need to \"png\" or \"jpg\" or \"gif\"")
		os.Exit(1)
	}

	wd, _ := os.Getwd()
	var saveFileName string
	ext := filepath.Ext(*filename)
	if ext != "."+*format {
		// filenameの拡張子を排除した部分
		n := filepath.Base((*filename)[:len(*filename)-len(filepath.Ext(*filename))])
		saveFileName = filepath.Join(wd, n+"."+*format)
	} else {
		saveFileName = filepath.Join(wd, *filename)
	}

	fmt.Println("output filename:", saveFileName)

	f, err := os.Create(saveFileName)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	if err := toJPEG(os.Stdin, f, *format); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer, ext string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format = ", kind)

	switch ext {
	case "png":
		return png.Encode(out, img)
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return errors.New("invalid encode extension")
	}
}
