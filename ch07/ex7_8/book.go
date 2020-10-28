package book

import "time"

type Book struct {
	Title     string
	Writer    string
	PublishAt time.Time
	Pages     int
}

func TitleCmp(b1, b2 *Book) int {
	if b1.Title < b2.Title {
		return 1
	}
	if b1.Title == b2.Title {
		return 0
	}
	return -1
}

func WriterCmp(b1, b2 *Book) int {
	if b1.Writer < b2.Writer {
		return 1
	}
	if b1.Writer == b2.Writer {
		return 0
	}
	return -1
}

func PublishCmp(b1, b2 *Book) int {
	if b1.PublishAt.Before(b2.PublishAt) {
		return 1
	}

	if b1.PublishAt == b2.PublishAt {
		return 0
	}

	return -1
}

func PageCmp(b1, b2 *Book) int {
	return b2.Pages - b1.Pages
}
