package command

import "github.com/hatobus/go-training/ch04/ex4_5/ex4.11/github"

func CloseIssue(owner, repo, number string) error {
	issue, err := github.EditIssue(owner, repo, number, map[string]string{"state": "close"})
	if err != nil {
		return err
	}
	github.PPIssueDetail(issue)
	return nil
}
