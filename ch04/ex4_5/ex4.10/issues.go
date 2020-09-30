package issues

import (
	"time"

	"golang.org/x/xerrors"
	"gopl.io/ch4/github"
)

type IssueTerm struct {
	PastDay   []*github.Issue
	PastMonth []*github.Issue
	PastYear  []*github.Issue
}

func GetIssueFromPeriod(terms []string) (*IssueTerm, error) {
	result, err := github.SearchIssues(terms)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	day := make([]*github.Issue, 0)
	month := make([]*github.Issue, 0)
	year := make([]*github.Issue, 0)

	yesterday := now.AddDate(0, 0, -1)
	previous_month := now.AddDate(0, -1, 0)
	previous_year := now.AddDate(-1, 0, 0)

	for _, issue := range result.Items {
		switch {
		case issue.CreatedAt.After(yesterday):
			day = append(day, issue)
		case issue.CreatedAt.After(previous_month):
			month = append(month, issue)
		case issue.CreatedAt.After(previous_year):
			year = append(year, issue)
		default:
			return nil, xerrors.New("invalid time")
		}
	}

	return &IssueTerm{
		PastDay:   day,
		PastMonth: month,
		PastYear:  year,
	}, nil
}
