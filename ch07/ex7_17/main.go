package xmlselect

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func SelectFromFile(filename string, selectElems []string) ([]string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("%v is not found", filename)
	} else if err != nil {
		return nil, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return XMLSelect(f, selectElems)
}

func XMLSelect(r io.Reader, selectElems []string) ([]string, error) {
	dec := xml.NewDecoder(r)
	var stk []string
	var attrs []map[string]string

	res := make([]string, 0)

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stk = append(stk, tok.Name.Local)
			attr := make(map[string]string, len(tok.Attr))
			for _, a := range tok.Attr {
				attr[a.Name.Local] = a.Value
			}
			attrs = append(attrs, attr)
		case xml.EndElement:
			stk = stk[:len(stk)-1]
			attrs = attrs[:len(attrs)-1]
		case xml.CharData:
			if containsAll(stringSlice(stk, attrs), selectElems) {
				res = append(res, string(tok))
			}
		}
	}
	return res, nil
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

// <div key="value> --> {"div", "key=value"}
func stringSlice(stk []string, attrs []map[string]string) []string {
	res := make([]string, 0)
	for i := range stk {
		res = append(res, stk[i])
		for key, val := range attrs[i] {
			res = append(res, fmt.Sprintf("%v=%v", key, val))
		}
	}
	return res
}
