package main

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var defaultColor = []string{"red", "blue"}

func getSVGHandler() http.HandlerFunc {
	return generateSVGHandler
}

func prepareSVGData(t testing.TB ,height, width int, colors []string) string {
	t.Helper()

	d := new(bytes.Buffer)
	err := genSVG(d, height, width, colors)
	if err != nil {
		t.Fatal(err)
	}

	return d.String()
}

func TestSVGServer(t *testing.T) {
	type testData struct {
		height string
		width string
		colors []string
		wantStatusCode int
		wantErr bool
		apiOut string
	}

	testCases := map[string]testData {
		"通常時(何も指定しない)": {
			height: "",
			width: "",
			colors: []string{},
			wantStatusCode: http.StatusCreated,
			apiOut: prepareSVGData(t, height, width, defaultColor),
		},
		"heightがマイナス": {
			height: "-1",
			width: "",
			colors: []string{},
			wantStatusCode: http.StatusBadRequest,
			wantErr: true,
			apiOut: "invalid height value\n",
		},
	}

	ts := httptest.NewServer(getSVGHandler())
	t.Cleanup(func() {
		ts.Close()
	})

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			var u bytes.Buffer
			u.WriteString(string(ts.URL))

			client := http.DefaultClient

			req, err := http.NewRequest(http.MethodGet, u.String(), nil)
			if err != nil {
				t.Fatal(err)
			}

			q := req.URL.Query()
			q.Add("height", tc.height)
			q.Add("width", tc.width)
			q.Add("color", strings.Join(tc.colors, ","))

			req.URL.RawQuery = q.Encode()

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			b, err := ioutil.ReadAll(resp.Body)
			bodyFromAPI := string(b)

			if resp.StatusCode != http.StatusCreated {
				if tc.wantErr == false {
					t.Fatalf("invalid response, error occured. err: %v", bodyFromAPI)
				} else if resp.StatusCode != tc.wantStatusCode {
					t.Fatalf("invalid status code, want %v but got %v, response %v", tc.wantStatusCode, resp.StatusCode, bodyFromAPI)
				} else {
					if diff := cmp.Diff(bodyFromAPI, tc.apiOut); diff != "" {
						t.Fatalf("mismatch api response, diff: %v", diff)
					}
				}
			} else {
				if diff := cmp.Diff(bodyFromAPI, tc.apiOut); diff != "" {
					t.Fatalf("mismatch api response, diff: %v", diff)
				}
			}
		})
	}
}
