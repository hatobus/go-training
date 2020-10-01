package github

import "fmt"

func PPIssues(issues []*Issue) {
	for _, i := range issues {
		PPIssue(i)
	}
}

func PPIssue(issue *Issue) {
	fmt.Printf("No: #%v\tUser: %v\tTitle: %v\n", issue.Number, issue.User.Login, issue.Title)
}

func PPIssueDetail(issue *Issue) {
	fmt.Printf(
		"No: #%v\tUser: %v\tTitle: %v\tState: %v\n"+
			"URL: %v\n"+"Body: %v\n",
		issue.Number, issue.User.Login, issue.Title, issue.State,
		issue.HTMLURL, issue.Body,
	)
	return
}
