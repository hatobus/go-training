package github

import (
	"fmt"
	"time"
)

var IssuesURL = "https://api.github.com/search/issues"
var APIURL = "https://api.github.com"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type RepoIssues struct {
	Issues       []Issue
	IssuesNumber map[int]Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

func (i Issue) GenIssueDetailURL() string {
	return fmt.Sprintf("/issue?no=%v", i.Number)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
