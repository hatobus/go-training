package command

import "github.com/hatobus/go-training/ch04/ex4_5/ex4.11/github"

func OpenIssue(owner, repo, number string) error {
	issue, err := github.EditIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		return err
	}
	github.PPIssueDetail(issue)
	return nil
}
