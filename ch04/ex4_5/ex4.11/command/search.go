package command

import (
	"github.com/hatobus/go-training/ch04/ex4_5/ex4.11/github"
)

func Search(args []string) error {
	result, err := github.SearchIssues(args)
	if err != nil {
		return err
	}
	github.PPIssues(result.Items)
	return nil
}
