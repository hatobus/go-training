package issues

import (
	"testing"
)

func TestGetIssueFromTerm(t *testing.T) {
	terms := []string{"repo:moby/moby", "is:open"}

	issue, err := GetIssueFromPeriod(terms)
	if err != nil {
		t.Fatal(err)
	}
}
