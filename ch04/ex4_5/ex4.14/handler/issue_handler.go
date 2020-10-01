package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/templates"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/github"
)

func NewIssuesHandler(issue *github.RepoIssues) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		if err := templates.TemplateOfIssueList.Execute(w, issue); err != nil {
			log.Println(err)
			http.Error(w, "failed to create HTML", http.StatusInternalServerError)
			return
		}
	}
}

func IssueDetailHandler(issue *github.RepoIssues) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := make(map[string]string)
		v := r.URL.Query()
		if v == nil {
			http.Error(w, "issue no not found", http.StatusBadRequest)
			return
		}

		for key, val := range v {
			params[key] = val[0]
		}

		issueNo, ok := params["no"]
		if !ok {
			http.Error(w, "issue no not found", http.StatusBadRequest)
			return
		}

		issueNoInt, err := strconv.Atoi(issueNo)
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid issue no", http.StatusInternalServerError)
			return
		}

		if err := templates.TemplateOfIssue.Execute(w, issue.IssuesNumber[issueNoInt]); err != nil {
			log.Println(err)
			http.Error(w, "failed to create HTML", http.StatusInternalServerError)
			return
		}
	}
}
