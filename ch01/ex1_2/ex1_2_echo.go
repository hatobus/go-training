package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	for i, arg := range os.Args[:] {
		s += sep + strconv.Itoa(i) + arg
		sep = " "
	}
	fmt.Printf(s)
}
