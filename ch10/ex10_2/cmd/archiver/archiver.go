package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hatobus/go-training/ch10/ex10_2/file_printer"
	_ "github.com/hatobus/go-training/ch10/ex10_2/tar"
	_ "github.com/hatobus/go-training/ch10/ex10_2/zip"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("invalid argument")
	}

	exitCode := 0
	for _, fileName := range os.Args[1:] {
		f, err := file_printer.FilePrint(fileName)
		if err != nil {
			fmt.Errorf("%v", err)
			exitCode = 2
		}
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			fmt.Errorf("%v", err)
			exitCode = 2
		}
	}

	os.Exit(exitCode)
}
