package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func fixtureServer(t testing.TB) *PriceDB {
	t.Helper()

	db := &PriceDB{}
	db.db = make(map[string]int, 0)
	db.db["shoe"] = 50
	db.db["socks"] = 5

	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Read)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)

	return db
}

func TestPriceDBCreate(t *testing.T) {
	//t.Parallel()
	db := fixtureServer(t)

	testCases := map[string]struct {
		requestURL       string
		expectStatusCode int
		expectRespBody   string
	}{
		"watchを作成できる": {
			requestURL:       "/create?item=watch&price=100",
			expectStatusCode: http.StatusOK,
			expectRespBody:   "",
		},
		"shoeはすでにあるので作成できない": {
			requestURL:       "/create?item=shoe&price=100",
			expectStatusCode: http.StatusBadRequest,
			expectRespBody:   "shoe is already exist\n",
		},
		"itemが指定されていない": {
			requestURL:       "/create",
			expectStatusCode: http.StatusBadRequest,
			expectRespBody:   "item is expected\n",
		},
		"priceが指定されていない": {
			requestURL:       "/create?item=shirts",
			expectStatusCode: http.StatusBadRequest,
			expectRespBody:   "price is expected\n",
		},
		"priceに不正な文字が含まれている": {
			requestURL:       "/create?item=hoge&price=fuga",
			expectStatusCode: http.StatusBadRequest,
			expectRespBody:   "price error invalid price input\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodPost, tc.requestURL, nil)
			rec := httptest.NewRecorder()

			db.ServeHTTP(rec, req)

			if diff := cmp.Diff(tc.expectStatusCode, rec.Code); diff != "" {
				t.Fatalf("invalid status code diff : %v", diff)
			}

			if diff := cmp.Diff(tc.expectRespBody, rec.Body.String()); diff != "" {
				t.Fatalf("invalid response body diff: %v", diff)
			}
		})
	}
}

func TestPriceDBRead(t *testing.T) {
	//t.Parallel()
	db := fixtureServer(t)

	testCases := map[string]struct {
		requestURL       string
		expectStatusCode int
		expectRespBody   string
	}{
		"shoeは存在する": {
			requestURL:       "/read?item=shoe",
			expectStatusCode: http.StatusOK,
			expectRespBody:   fmt.Sprintf("%v: %v\n", "shoe", 50),
		},
		"watchは存在しない": {
			requestURL:       "/read?item=watch",
			expectStatusCode: http.StatusNotFound,
			expectRespBody:   fmt.Sprintf("%v not found\n", "watch"),
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, tc.requestURL, nil)
			rec := httptest.NewRecorder()

			db.ServeHTTP(rec, req)

			if diff := cmp.Diff(tc.expectStatusCode, rec.Code); diff != "" {
				t.Fatalf("invalid status code diff : %v", diff)
			}

			if diff := cmp.Diff(tc.expectRespBody, rec.Body.String()); diff != "" {
				t.Fatalf("invalid response body diff: %v", diff)
			}
		})
	}
}

func TestPriceDBUpdate(t *testing.T) {
	//t.Parallel()
	db := fixtureServer(t)

	testCases := map[string]struct {
		requestURL         string
		expectStatusCode   int
		expectRespBody     string
		appendixAssertFunc func(testing.TB)
	}{
		"shoeはupdateできる": {
			requestURL:       "/read?item=shoe&price=100",
			expectStatusCode: http.StatusOK,
			expectRespBody:   "",
			appendixAssertFunc: func(t testing.TB) {
				t.Helper()

				db.m.RLock()
				defer db.m.RUnlock()
				if _, ok := db.db["shoe"]; !ok {
					t.Fatalf("shoe is not found")
				}

				shoePrice, _ := db.db["shoe"]
				if diff := cmp.Diff(shoePrice, 100); diff != "" {
					t.Fatalf("invalid new price diff = %v", diff)
				}
			},
		},
		"watchは存在しない": {
			requestURL:       "/read?item=watch&price=10",
			expectStatusCode: http.StatusNotFound,
			expectRespBody:   "watch is not found\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodPut, tc.requestURL, nil)
			rec := httptest.NewRecorder()

			db.ServeHTTP(rec, req)

			if diff := cmp.Diff(tc.expectStatusCode, rec.Code); diff != "" {
				t.Fatalf("invalid status code diff : %v", diff)
			}

			if diff := cmp.Diff(tc.expectRespBody, rec.Body.String()); diff != "" {
				t.Fatalf("invalid response body diff: %v", diff)
			}

			if tc.appendixAssertFunc != nil {
				tc.appendixAssertFunc(t)
			}
		})
	}
}

func TestPriceDBDelete(t *testing.T) {
	//t.Parallel()
	db := fixtureServer(t)

	testCases := map[string]struct {
		requestURL         string
		expectStatusCode   int
		expectRespBody     string
		appendixAssertFunc func(testing.TB)
	}{
		"shoeはdeleteできる": {
			requestURL:       "/delete?item=shoe",
			expectStatusCode: http.StatusOK,
			expectRespBody:   "",
			appendixAssertFunc: func(t testing.TB) {
				t.Helper()

				db.m.RLock()
				defer db.m.RUnlock()
				if _, ok := db.db["shoe"]; ok {
					t.Fatalf("shoe is exist")
				}

			},
		},
		"watchは存在しない": {
			requestURL:       "/delete?item=watch",
			expectStatusCode: http.StatusNotFound,
			expectRespBody:   "watch is not found\n",
		},
		"itemを指定しない": {
			requestURL:       "/delete",
			expectStatusCode: http.StatusBadRequest,
			expectRespBody:   "item is expected\n",
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodDelete, tc.requestURL, nil)
			rec := httptest.NewRecorder()

			db.ServeHTTP(rec, req)

			if diff := cmp.Diff(tc.expectStatusCode, rec.Code); diff != "" {
				t.Fatalf("invalid status code diff : %v", diff)
			}

			if diff := cmp.Diff(tc.expectRespBody, rec.Body.String()); diff != "" {
				t.Fatalf("invalid response body diff: %v", diff)
			}

			if tc.appendixAssertFunc != nil {
				tc.appendixAssertFunc(t)
			}
		})
	}
}
