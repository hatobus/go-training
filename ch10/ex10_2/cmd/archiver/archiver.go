package main

import (
	"fmt"
	"os"

	"github.com/hatobus/go-training/ch10/ex10_2/file_printer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("invalid argument")
	}

	exitCode := 0
	for _, fileName := range os.Args[1:] {
		err := file_printer.FilePrint(fileName)
		if err != nil {
			fmt.Errorf("%v", err)
			exitCode = 2
		}
	}

	os.Exit(exitCode)
}
