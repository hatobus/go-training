package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/github"
	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/handler"
)

const usage = `
	usage: githubIssueServer owner repo
`

func main() {
	if len(os.Args) != 3 {
		log.Fatal(usage)
	}

	owner := os.Args[1]
	repo := os.Args[2]

	issues, err := github.ReadIssueFromIdentifier(owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	// Goのルーティングは最初にマッチしたものなので / にするとそれ以外のエンドポイントが作れない
	http.Handle("/list", handler.NewIssuesHandler(issues))
	http.Handle("/issue", handler.IssueDetailHandler(issues))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
