package book_sort

import "github.com/hatobus/go-training/ch07/ex7_8/book"

const (
	lt = iota - 1
	eq
	gt
)

type BookSort struct {
	books  []*book.Book
	cmpFun []func(b1 *book.Book, b2 *book.Book) int
}

func NewBookSorter(books []*book.Book, compareFunc []func(*book.Book, *book.Book) int) BookSort {
	bs := BookSort{}
	bs.books = books
	bs.cmpFun = append(bs.cmpFun, compareFunc...)
	return bs
}

func (bs BookSort) Len() int {
	return len(bs.books)
}

func (bs BookSort) Less(i, j int) bool {
	for _, comparator := range bs.cmpFun {
		c := comparator(bs.books[i], bs.books[j])
		switch c {
		case lt:
			return true
		case eq:
			continue
		case gt:
			return false
		}
	}
	return false
}

func (bs BookSort) Swap(i, j int) { bs.books[i], bs.books[j] = bs.books[j], bs.books[i] }
