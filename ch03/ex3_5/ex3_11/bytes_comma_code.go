package comma

import (
	"bytes"
	"strings"

	"golang.org/x/xerrors"
)

func BytesCommaWithCode(s string) (string, error) {
	buf := new(bytes.Buffer)

	elems := strings.Split(s, ".")
	if len(elems) > 2 {
		return "", xerrors.Errorf("invalid format")
	}

	var upper, lower string
	if len(elems) == 1 {
		upper = elems[0]
		lower = ""
	} else {
		upper = elems[0]
		lower = elems[1]
	}

	if upper == "" {
		return "", xerrors.Errorf("invalid format")
	} else if len(upper) <= 3 {
		return s, nil
	}

	if lower != "" {
		lower = "." + lower
	}

	buf.WriteString(upperComma(upper))
	buf.WriteString(lower)

	return buf.String(), nil
}

func upperComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	buf := new(bytes.Buffer)

	runes := []rune(s)

	if s[0] == '-' || s[0] == '+' {
		buf.WriteRune(runes[0])
		runes = runes[1:]
	}

	start := len(runes) % 3
	if start == 0 {
		start = 3
	}

	buf.WriteString(string(runes[:start]))
	for i := start; i < len(runes); i += 3 {
		buf.WriteString(",")
		buf.WriteString(string(runes[i : i+3]))
	}

	return buf.String()
}
