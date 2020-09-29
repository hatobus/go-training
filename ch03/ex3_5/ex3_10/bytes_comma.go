package comma

import "bytes"

func BytesComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	buf := new(bytes.Buffer)

	start := n % 3
	if start == 0 {
		start = 3
	}

	buf.WriteString(s[:start])

	for i := start; i < len(s); i += 3 {
		buf.WriteString(",")
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
