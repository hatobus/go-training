package main

import (
	"io/ioutil"
	"net/http"
)

const list_api_response = "./testData/raw_text_all_issue.txt"

func main() {
	http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadFile(list_api_response)
		if err != nil {
			http.Error(w, "failed to read test data", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(bytes)
		return
	})
}
