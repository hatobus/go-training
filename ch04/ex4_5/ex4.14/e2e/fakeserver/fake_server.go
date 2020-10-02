package fakeserver

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-cmp/cmp"
)

const list_api_response = "./testData/raw_text_all_issue.txt"

func FakeIssueListAPIHandler(requestParamPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if diff := cmp.Diff(requestParamPath, r.URL.RawPath); diff != "" {
			http.Error(w, fmt.Sprintf("mismatch API Handler URL, diff: %v", diff), http.StatusBadRequest)
		}
		bytes, err := ioutil.ReadFile(list_api_response)
		if err != nil {
			http.Error(w, "failed to read test data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	}
}
