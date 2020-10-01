package main

import (
	"log"
	"os"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.13/command"
)

const (
	cmdGet = "get"
)

const usage = `
poster: get poster image if you want
	poster get {name}: get poster image you want
`

func main() {
	if len(os.Args) < 2 {
		log.Fatal(usage)
	}

	commandStr := os.Args[1]
	switch commandStr {
	case cmdGet:
		title := os.Args[2]
		poster, err := command.GetMovie(title)
		if err != nil {
			log.Printf("failed to get %v's poster\n", title)
			log.Print(err)
			return
		}
		log.Printf("poster correctly saved: %v\n", poster)
	default:
		log.Fatal(usage)
	}
}
