package github

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path"

	"golang.org/x/xerrors"
)

func EditIssue(owner, repo, number string, jsonFields map[string]string) (*Issue, error) {
	// map[string]string はそのまんま json.Marshal ができる
	b, err := json.Marshal(jsonFields)

	client := &http.Client{}
	u, err := url.Parse(APIURL)
	if err != nil {
		return nil, err
	}

	u.Path = path.Join(u.Path, "repos", owner, repo, "issues", number)
	req, err := http.NewRequest(http.MethodPatch, u.String(), bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(os.Getenv("GITHUB_USERNAME"), os.Getenv("GITHUB_PASSWD"))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, xerrors.Errorf("faild to edit issue status code: %v, status: %v", resp.StatusCode, resp.Status)
	}

	var issue Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}

	return &issue, nil
}
