package main

import (
	"log"
	"os"

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
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	mode := os.Args[1]
	args := os.Args[2:]

	if mode == cmdSearch {
		if len(args) < 1 {
			log.Fatalf(usage)
		}
		if err := command.Search(args); err != nil {
			log.Fatal(err)
		}
		return
	}
}
