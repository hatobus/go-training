package ex5_2

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func HTMLTagFrequency(htmlFileName string) (map[string]int, error) {
	_, err := os.Stat(htmlFileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("%v is not found", htmlFileName)
	} else if err != nil {
		return nil, err
	}

	f, err := os.Open(htmlFileName)
	if err != nil {
		return nil, err
	}

	tags := make(map[string]int)

	htmlData := html.NewTokenizer(f)

	for {
		t := htmlData.Next()
		if t == html.ErrorToken {
			break
		}
		tagName, _ := htmlData.TagName()
		if string(tagName) != "" {
			tags[string(tagName)]++
		}
	}

	return tags, nil
}
