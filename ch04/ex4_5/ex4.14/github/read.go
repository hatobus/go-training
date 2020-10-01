package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

func ReadIssueFromIdentifier(owner, repo string) ([]Issue, error) {
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

	return issues, nil
}
