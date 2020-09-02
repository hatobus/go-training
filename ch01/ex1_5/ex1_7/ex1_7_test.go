package main

import (
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"os/exec"
	"testing"
)

func TestIoCopy(t *testing.T) {
	type testData struct {
		URLs []string
		ExpectError bool
		ExpectOutputs []string
		ExpectErrorString string
	}

	testCases := map[string]testData{
		"gopl.ioから取得が出来る": {
			URLs: []string{"http://gopl.io"},
			ExpectError: false,
			ExpectOutputs: []string{"../gopl.html"},
		},
		"存在しないURLの場合": {
			URLs: []string{"http://bad.gopl.io"},
			ExpectError: true,
			ExpectErrorString: "exit status 1",
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T){
			args := make([]string, 0, len(tc.URLs)+2)

			args = append(args, "run", "ex1_7.go")
			args = append(args, tc.URLs...)

			out, err := exec.Command("go", args...).Output()
			if err != nil {
				if !tc.ExpectError || err.Error() != tc.ExpectErrorString {
					t.Fatalf("failed to execute ex1_7.go: err: %v, output: %v", err, string(out))
				}
			} else {
				if tc.ExpectError {
					t.Fatalf("unexpected output, expected error but not occured")
				}

				var expectOuts string
				for _, html := range tc.ExpectOutputs {
					bytes, err := ioutil.ReadFile(html)
					if err != nil {
						t.Fatalf("compiration file read failed: %v", err)
					}
					expectOuts += string(bytes)
				}

				if diff := cmp.Diff(string(out), expectOuts); diff != "" {
					t.Log(string(out))
					t.Fatalf("fetched data and expected data unmmach: %v", diff)
				}
			}
		})
	}
}
