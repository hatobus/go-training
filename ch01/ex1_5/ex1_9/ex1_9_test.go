package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/hatobus/go-training/pioutil"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestFetchWithStatusCode(t *testing.T) {
	type testData struct {
		URLs []string
		ExpectStatusCode int
		ExpectError bool
		ExpectOutputs []string
		ExpectErrorString string
	}

	testCases := map[string]testData{
		"gopl.ioから取得が出来てステータスコードが200": {
			URLs: []string{"http://gopl.io"},
			ExpectStatusCode: http.StatusOK,
			ExpectError: false,
			ExpectOutputs: []string{"../gopl.html"},
		},
		"bad.gopl.ioから取得が出来ずステータスコードは404": {
			URLs: []string{"bad.gopl.io"},
			ExpectStatusCode: http.StatusNotFound,
			ExpectError: false,
			ExpectOutputs: []string{"../gopl.html"},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){
			os.Args = []string{"ex1_9.go"}
			os.Args = append(os.Args, tc.URLs...)

			out := pioutil.OutputCapture(func(){
				main()
			})
			if tc.ExpectError{
				if string(out) != tc.ExpectErrorString {
					t.Fatalf("failed to execute ex1_9.go: output: %v", out)
				}
			} else {
				if tc.ExpectError {
					t.Fatalf("unexpected output, expected error but not occured: %v", out)
				}

				var expectOuts string
				for i := 0; i < len(tc.URLs); i++ {
					u := tc.URLs[i]
					if !strings.HasPrefix(u, "http://") {
						u = "http://" + u
					}
					expectOuts += fmt.Sprintf("%v response: %v\n", u, tc.ExpectStatusCode)
					bytes, err := ioutil.ReadFile(tc.ExpectOutputs[i])
					if err != nil {
						t.Fatalf("compiration file read failed: %v", err)
					}
					expectOuts += string(bytes)
				}

				if diff := cmp.Diff(out, expectOuts); diff != "" {
					t.Log(out)
					t.Log(expectOuts)
					t.Fatalf("fetched data and expected data unmmach: %v", diff)
				}
			}
		})
	}

}

