package book

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

func fixtureBook1(t testing.TB, fs ...func(*Book)) *Book {
	t.Helper()

	baseBook1 := Book{
		Title:     "Hear The Wind Sing",
		Writer:    "Murakami Haruki",
		PublishAt: time.Date(1979, 7, 23, 0, 0, 0, 0, jst),
		Pages:     201,
	}

	for _, f := range fs {
		f(&baseBook1)
	}

	return &baseBook1
}

func fixtureBook2(t testing.TB, fs ...func(*Book)) *Book {
	t.Helper()

	baseBook2 := Book{
		Title:     "Almost Transparent Blue",
		Writer:    "Murakami Ryu",
		PublishAt: time.Date(2004, 2, 1, 0, 0, 0, 0, jst),
		Pages:     209,
	}

	for _, f := range fs {
		f(&baseBook2)
	}

	return &baseBook2
}

func TestTitleCmp(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		b1        *Book
		b2        *Book
		expectOut int
	}{
		"b1 < b2のタイトル": {
			b1:        fixtureBook1(t, func(*Book) {}),
			b2:        fixtureBook2(t, func(*Book) {}),
			expectOut: -1,
		},
		"b1 == b2のタイトル": {
			b1: fixtureBook1(t, func(b *Book) {
				b.Title = "same title"
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.Title = "same title"
			}),
			expectOut: 0,
		},
		"b1 > b2のタイトル": {
			b1: fixtureBook1(t, func(b *Book) {
				b.Title = "a-book"
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.Title = "z-book"
			}),
			expectOut: 1,
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := TitleCmp(tc.b1, tc.b2)

			if diff := cmp.Diff(out, tc.expectOut); diff != "" {
				t.Fatalf("ivalid output diff: %v", diff)
			}
		})
	}
}

func TestWriterCmp(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		b1        *Book
		b2        *Book
		expectOut int
	}{
		"b1 < b2のWriter": {
			b1:        fixtureBook1(t, func(*Book) {}),
			b2:        fixtureBook2(t, func(*Book) {}),
			expectOut: 1,
		},
		"b1 == b2のWriter": {
			b1: fixtureBook1(t, func(b *Book) {
				b.Writer = "same title"
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.Writer = "same title"
			}),
			expectOut: 0,
		},
		"b1 > b2のWriter": {
			b1: fixtureBook1(t, func(b *Book) {
				b.Writer = "z-writer"
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.Writer = "a-writer"
			}),
			expectOut: -1,
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := WriterCmp(tc.b1, tc.b2)

			if diff := cmp.Diff(out, tc.expectOut); diff != "" {
				t.Fatalf("ivalid output diff: %v", diff)
			}
		})
	}
}

func TestPublishCmp(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		b1        *Book
		b2        *Book
		expectOut int
	}{
		"b1 < b2のPublishDate": {
			b1:        fixtureBook1(t, func(*Book) {}),
			b2:        fixtureBook2(t, func(*Book) {}),
			expectOut: 1,
		},
		"b1 == b2のPublishDate": {
			b1: fixtureBook1(t, func(b *Book) {
				b.PublishAt = time.Date(2006, 1, 2, 3, 4, 5, 6, jst)
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.PublishAt = time.Date(2006, 1, 2, 3, 4, 5, 6, jst)
			}),
			expectOut: 0,
		},
		"b1 > b2のPublishAt": {
			b1: fixtureBook1(t, func(b *Book) {
				b.PublishAt = time.Date(2001, 1, 1, 0, 0, 0, 0, jst)
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.PublishAt = time.Date(2000, 1, 1, 0, 0, 0, 0, jst)
			}),
			expectOut: -1,
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := PublishCmp(tc.b1, tc.b2)

			if diff := cmp.Diff(out, tc.expectOut); diff != "" {
				t.Fatalf("ivalid output diff: %v", diff)
			}
		})
	}
}

func TestPageCmp(t *testing.T) {
	t.Parallel()

	testData := map[string]struct {
		b1         *Book
		b2         *Book
		assertFunc func(testing.TB, int)
	}{
		"b1 < b2のPage": {
			b1: fixtureBook1(t, func(*Book) {}),
			b2: fixtureBook2(t, func(*Book) {}),
			assertFunc: func(t testing.TB, pages int) {
				t.Helper()
				if pages <= 0 {
					t.Fatalf("invalid output want less than 0, but >= 0, pages: %v", pages)
				}
			},
		},
		"b1 == b2のPages": {
			b1: fixtureBook1(t, func(b *Book) {
				b.Pages = 200
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.Pages = 200
			}),
			assertFunc: func(t testing.TB, pages int) {
				t.Helper()
				if pages != 0 {
					t.Fatalf("invalid output want 0, but not 0, pages: %v", pages)
				}
			},
		},
		"b1 > b2のPage": {
			b1: fixtureBook1(t, func(b *Book) {
				b.Pages = 100
			}),
			b2: fixtureBook2(t, func(b *Book) {
				b.Pages = 200
			}),
			assertFunc: func(t testing.TB, pages int) {
				t.Helper()
				if pages <= 0 {
					t.Fatalf("invalid output want larger than 0, but <= 0, pages: %v", pages)
				}
			},
		},
	}

	for testName, tc := range testData {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			out := PageCmp(tc.b1, tc.b2)

			tc.assertFunc(t, out)
		})
	}
}
