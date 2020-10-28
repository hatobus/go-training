package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/hatobus/go-training/ch07/ex7_8/book"
	booksort "github.com/hatobus/go-training/ch07/ex7_8/book_sort"
)

var tz *time.Location

var books []*book.Book

func init() {
	var err error
	tz, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	books = []*book.Book{
		{
			Title:     "Life for Sale",
			Writer:    "Mishima Yuikio",
			PublishAt: time.Date(1968, 12, 25, 0, 0, 0, 0, tz),
			Pages:     261,
		},
		{
			Title:     "Oh father",
			Writer:    "Isaka Kotaro",
			PublishAt: time.Date(2012, 7, 1, 0, 0, 0, 0, tz),
			Pages:     542,
		},
		{
			Title:     "Lemon",
			Writer:    "Kajii Motojiro",
			PublishAt: time.Date(1925, 1, 1, 0, 0, 0, 0, tz),
			Pages:     9,
		},
		{
			Title:     "Sputnik Sweetheart",
			Writer:    "Murakami Haruki",
			PublishAt: time.Date(2001, 4, 15, 0, 0, 0, 0, tz),
			Pages:     318,
		},
	}
}

var tpl = template.Must(template.New("book").Parse(`
<html>
<body>

<table>
	<tr>
		<th><a href="?booksort=title">title</a></th>
		<th><a href="?booksort=writer">writer</a></th>
		<th><a href="?booksort=publishat">publishat</a></th>
		<th><a href=">booksort=pages">pages</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Title}}</td>
		<td>{{.Writer}}</td>
		<td>{{.PublishAt}}</td>
		<td>{{.Pages}} pages</td>
	</tr>
{{end}}
</body>
</html>
`))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var b booksort.BookSort
		switch r.FormValue("booksort") {
		case "title":
			b = booksort.NewBookSorter(books, []func(*book.Book, *book.Book) int{book.TitleCmp})
		case "writer":
			b = booksort.NewBookSorter(books, []func(*book.Book, *book.Book) int{book.WriterCmp})
		case "publishat":
			b = booksort.NewBookSorter(books, []func(*book.Book, *book.Book) int{book.PublishCmp})
		case "pages":
			b = booksort.NewBookSorter(books, []func(*book.Book, *book.Book) int{book.PageCmp})
		}
		sort.Sort(b)
		err := tpl.Execute(w, books)
		if err != nil {
			log.Println(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
