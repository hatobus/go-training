package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCalclatorServer(t *testing.T) {
	testCases := map[string]struct {
		expression string
		variables  map[string]string
		answer     float64
	}{
		"x+y": {
			expression: "x+y",
			variables: map[string]string{
				"x": "1", "y": "2",
			},
			answer: 3,
		},
		"log2(x)-log10(y)": {
			expression: "log2(x)-log10(y)",
			variables: map[string]string{
				"x": "8", "y": "1000",
			},
			answer: 0,
		},
	}

	ts := httptest.NewServer(getCalclatorHandler())
	t.Cleanup(func() {
		ts.Close()
	})

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			u, err := url.Parse(ts.URL)
			if err != nil {
				t.Fatal(err)
			}

			q := u.Query()

			q.Set("expression", tc.expression)

			for k, v := range tc.variables {
				q.Set(k, v)
			}

			u.RawQuery = q.Encode()

			res, err := http.Get(u.String())
			if err != nil {
				t.Fatal(err)
			}
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}

			var ans Answer
			err = json.Unmarshal(b, &ans)
			if err != nil {
				t.Fatal(string(b))
				t.Fatal(err)
			}

			if diff := cmp.Diff(ans.Answer, tc.answer); diff != "" {
				t.Fatalf("invalid answer diff: %v", diff)
			}
		})
	}
}
