package cmd

import (
	"os"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/cmd/main"
)

func RunMain(args []string) {
	os.Args = args
	main.Main()
}
