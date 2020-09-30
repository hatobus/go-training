package command

import "github.com/hatobus/go-training/ch04/ex4_5/ex4.11/github"

func ReadIssue(owner, repo, number string) error {
	issue, err := github.ReadIssueFromIdentifier(owner, repo, number)
	if err != nil {
		return err
	}
	github.PPIssueDetail(issue)
	return nil
}
