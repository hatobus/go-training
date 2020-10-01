package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.12/argument"
	"github.com/hatobus/go-training/ch04/ex4_5/ex4.12/command"
)

const (
	cmdGet    = "get"
	cmdIndex  = "index"
	cmdSearch = "search"
)

var usage = `
xkcd command
xkcd get N: get N's story
xkcd index FILENAME: generate 
xkcd search INDEX_FILE QUERY: search index from query
`

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	barb := os.Args[1]
	switch barb {
	case cmdGet:
		if !argument.ValidateGetArguments(os.Args) {
			log.Fatal(usage)
		}
		c, err := command.GetComic(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(c)
		return
	case cmdIndex:
		if !argument.ValidateIndexArguments(os.Args) {
			log.Fatal(usage)
		}
		err := command.GenerateIndex(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		return
	case cmdSearch:
	default:
		log.Fatal(usage)
	}
}
