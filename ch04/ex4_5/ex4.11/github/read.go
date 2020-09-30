package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

func ReadIssueFromIdentifier(owner, repo, number string) (*Issue, error) {
	u, err := url.Parse(APIURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "repos", owner, repo, "issues", number)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
