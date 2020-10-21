package ex5_9

import (
	"regexp"
	"strings"
)

var pattern = regexp.MustCompile(`\$\w+`)

func Expand(s string, f func(string) string) string {
	return pattern.ReplaceAllStringFunc(s, func(s string) string {
		var dst string
		if strings.HasPrefix(s, "$") {
			dst = s[1:]
		} else {
			dst = s
		}
		return f(dst)
	})
}
