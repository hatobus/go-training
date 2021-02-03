package ex12_9

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

var (
	kindListStart = '('
	kindListEnd   = ')'
)

type Token interface{}

type Symbol struct{}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct{}
type EndList struct{}

type Decoder struct {
	scan  scanner.Scanner
	err   error
	depth int
}

func NewDecoder(r io.Reader) *Decoder {
	var scan scanner.Scanner
	scan.Init(r)

	return &Decoder{
		scan: scan,
	}
}

func (d *Decoder) Token() (Token, error) {
	token := d.scan.Scan()
	if d.err != nil {
		return nil, d.err
	}
	if d.depth == 0 && token != kindListStart && token != scanner.EOF {
		return nil, fmt.Errorf("unexpect character expect ( but fot %v", scanner.TokenString(token))
	}

	switch token {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		return d.scan.TokenText(), nil
	case scanner.String:
		text := d.scan.TokenText()
		// Assume all strings are quoted.
		return String{text[1 : len(text)-1]}, nil
	case scanner.Int:
		n, err := strconv.ParseInt(d.scan.TokenText(), 10, 64)
		if err != nil {
			return nil, err
		}
		return Int{int(n)}, nil
	case kindListStart:
		d.depth++
		return StartList{}, nil
	case kindListEnd:
		d.depth--
		return EndList{}, nil
	default:
		pos := d.scan.Pos()
		return nil, fmt.Errorf("unexpected token %s at L%d:C%d", scanner.TokenString(token), pos.Line, pos.Column)
	}
}
