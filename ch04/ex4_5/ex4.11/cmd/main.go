package main

import (
	"log"
	"os"

	arguments "github.com/hatobus/go-training/ch04/ex4_5/ex4.11/argument"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.11/command"
)

const (
	cmdSearch = "search"
	cmdRead   = "read"
	cmdEdit   = "edit"
	cmdClose  = "close"
	cmdOpen   = "open"
)

var usage string = `
usage:
	search: search issue. args --> search query
	[read | edit | close | open]: args --> owner repo issue_number
`

func main() {
	if !arguments.ValidateArgsRunning(os.Args) {
		log.Fatal(usage)
	}

	mode := os.Args[1]
	args := os.Args[2:]

	if mode == cmdSearch {
		if !arguments.ValidateSearchArguments(args) {
			log.Fatalf(usage)
		}
		if err := command.Search(args); err != nil {
			log.Fatal(err)
		}
		return
	}
}
