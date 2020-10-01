package github

import (
	"fmt"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

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

func (i Issue) GenURL() string {
	return fmt.Sprintf("/issues/%v", i.Number)
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
