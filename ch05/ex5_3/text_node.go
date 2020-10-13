package ex5_3

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func HTMLTextNodes(htmlFileName string) (map[string][]string, error) {
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

	textNodes := make(map[string][]string)

	htmlData := html.NewTokenizer(f)

	var lastTag string
	for {
		t := htmlData.Next()
		if t == html.ErrorToken {
			break
		} else if t == html.StartTagToken {
			tagName, _ := htmlData.TagName()
			lastTag = string(tagName)
		} else if t == html.TextToken {
			if lastTag == "script" || lastTag == "style" {
				continue
			}
			text := htmlData.Text()
			trimed := strings.TrimSpace(string(text))
			if len(trimed) != 0 {
				textNodes[lastTag] = append(textNodes[lastTag], trimed)
			}
		} else {
			continue
		}
	}

	return textNodes, nil
}
