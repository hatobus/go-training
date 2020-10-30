package book_sort

import (
	"sort"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/hatobus/go-training/ch07/ex7_8/book"
)

var jst *time.Location

var bookmap map[int]*book.Book

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	bookmap = map[int]*book.Book{
		0: {
			Title:     "Life for Sale",
			Writer:    "Mishima Yuikio",
			PublishAt: time.Date(1968, 12, 25, 0, 0, 0, 0, jst),
			Pages:     261,
		},
		1: {
			Title:     "Oh father",
			Writer:    "Isaka Kotaro",
			PublishAt: time.Date(2012, 7, 1, 0, 0, 0, 0, jst),
			Pages:     542,
		},
		2: {
			Title:     "Lemon",
			Writer:    "Kajii Motojiro",
			PublishAt: time.Date(1925, 1, 1, 0, 0, 0, 0, jst),
			Pages:     9,
		},
		3: {
			Title:     "Sputnik Sweetheart",
			Writer:    "Murakami Haruki",
			PublishAt: time.Date(2001, 4, 15, 0, 0, 0, 0, jst),
			Pages:     318,
		},
	}
}

// get the random books
func getBooks(t testing.TB) []*book.Book {
	t.Helper()
	bs := make([]*book.Book, 0, len(bookmap))
	for _, b := range bookmap {
		bs = append(bs, b)
	}
	return bs
}

func genExpectBooks(t testing.TB, bookids ...int) []*book.Book {
	t.Helper()

	if len(bookmap) != len(bookids) {
		t.Fatalf("invalid argument length")
	}

	b := make([]*book.Book, 0, len(bookids))
	for _, id := range bookids {
		b = append(b, bookmap[id])
	}

	return b
}

func TestBookSorter(t *testing.T) {
	testCases := map[string]struct {
		books       []*book.Book
		compareFunc []func(*book.Book, *book.Book) int
		expectBooks []*book.Book
	}{
		"Title順にSort": {
			books:       getBooks(t),
			compareFunc: []func(*book.Book, *book.Book) int{book.TitleCmp},
			expectBooks: genExpectBooks(t, 3, 1, 0, 2),
		},
		"Writer順にSort": {
			books:       getBooks(t),
			compareFunc: []func(*book.Book, *book.Book) int{book.WriterCmp},
			expectBooks: genExpectBooks(t, 3, 0, 2, 1),
		},
		"Publish順にSort": {
			books:       getBooks(t),
			compareFunc: []func(*book.Book, *book.Book) int{book.PublishCmp},
			expectBooks: genExpectBooks(t, 1, 3, 0, 2),
		},
		"Title順にSortしてPublish順にSort": {
			books:       getBooks(t),
			compareFunc: []func(*book.Book, *book.Book) int{book.TitleCmp, book.PublishCmp},
			expectBooks: genExpectBooks(t, 3, 1, 0, 2),
		},
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			sort.Sort(NewBookSorter(tc.books, tc.compareFunc))

			if diff := cmp.Diff(&tc.books, &tc.expectBooks); diff != "" {
				t.Fatalf("sort failed diff: %v", diff)
			}
		})
	}
}
