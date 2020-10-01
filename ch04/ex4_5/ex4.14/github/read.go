package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

func ReadIssueFromIdentifier(owner, repo string) (*RepoIssues, error) {
	u, err := url.Parse(APIURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "repos", owner, repo, "issues")

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}

	ri := &RepoIssues{}
	ri.IssuesNumber = make(map[int]Issue, len(issues))
	for _, issue := range issues {
		ri.Issues = append(ri.Issues, issue)
		ri.IssuesNumber[issue.Number] = issue
	}

	return ri, nil
}
