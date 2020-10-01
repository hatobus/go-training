package handler

import (
	"log"
	"net/http"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/templates"

	"github.com/hatobus/go-training/ch04/ex4_5/ex4.14/github"
)

func NeeIssuesHandler(issue []github.Issue) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if err := templates.TemplateOfIssue.Execute(w, issue[0]); err != nil {
			log.Println(err)
			http.Error(w, "failed to create HTML", http.StatusInternalServerError)
			return
		}
	}
}
