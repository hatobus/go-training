package main_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	ex34 "github.com/hatobus/go-training/ch03/ex3_2/ex3.4"
)

var defaultColor = []string{"red", "blue"}

func getSVGHandler() http.HandlerFunc {
	return ex34.GenerateSVGHandler
}

func prepareSVGData(t testing.TB, height, width int, colors []string) string {
	t.Helper()

	d := new(bytes.Buffer)
	err := ex34.GenSVG(d, height, width, colors)
	if err != nil {
		t.Fatal(err)
	}

	return d.String()
}

func TestSVGServer(t *testing.T) {
	type testData struct {
		height         string
		width          string
		colors         []string
		wantStatusCode int
		wantErr        bool
		apiOut         string
	}

	testCases := map[string]testData{
		"通常時(何も指定しない)": {
			height:         "",
			width:          "",
			colors:         []string{},
			wantStatusCode: http.StatusCreated,
			apiOut:         prepareSVGData(t, ex34.Height, ex34.Width, defaultColor),
		},
		"height 300 width 800": {
			height:         "300",
			width:          "800",
			colors:         []string{},
			wantStatusCode: http.StatusCreated,
			apiOut:         prepareSVGData(t, 300, 800, defaultColor),
		},
		"colorを指定する": {
			height:         "",
			width:          "",
			colors:         []string{"yellow", "green"},
			wantStatusCode: http.StatusCreated,
			apiOut:         prepareSVGData(t, ex34.Height, ex34.Width, []string{"yellow", "green"}),
		},
		"heightがマイナス": {
			height:         "-1",
			width:          "",
			colors:         []string{},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
			apiOut:         "invalid height value\n",
		},
		"widthがマイナス": {
			height:         "",
			width:          "-1",
			colors:         []string{},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
			apiOut:         "invalid width value\n",
		},
		"colorが多い": {
			height:         "",
			width:          "",
			colors:         []string{"blue", "yellow", "green"},
			wantStatusCode: http.StatusBadRequest,
			wantErr:        true,
			apiOut:         "invalid color, length must be 2\n",
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
