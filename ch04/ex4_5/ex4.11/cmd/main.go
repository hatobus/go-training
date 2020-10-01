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

	cmdEditor = "editor"
)

var usage string = `
usage:
	search: search issue. args --> search query
	[read | edit | close | open]: args --> owner repo issue_number
	editor: set issue editor --> editor [editor command]
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
	} else if mode == cmdEditor {
		if !arguments.ValidateEditorArguments(args) {
			log.Fatalf("editor %v is not valid", args[0])
		}
		err := os.Setenv("ISSUE_EDITOR", args[0])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current issue editor: %v\n", os.Getenv("ISSUE_EDITOR"))
		return
	}

	owner, repo, number, err := arguments.GetIdentifier(args)
	if err != nil {
		log.Println(err)
		log.Fatal(usage)
	}

	switch mode {
	case cmdRead:
		err = command.ReadIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
	case cmdOpen:
		err = command.OpenIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
	case cmdClose:
		err = command.CloseIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
	case cmdEdit:
		err = command.EditIssue(owner, repo, number)
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal(usage)
	}
}
