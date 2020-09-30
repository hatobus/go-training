package github

import "fmt"

func PPIssue(issues []*Issue) {
	for _, i := range issues {
		fmt.Printf("No: #%v\tUser: %v\tTitle: %v\n", i.Number, i.User.Login, i.Title)
	}
}
