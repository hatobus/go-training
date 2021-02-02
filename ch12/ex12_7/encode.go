package ex12_7

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"text/scanner"
)

type Decoder struct {
	lex *lexer
}

func NewDecoder(r io.Reader) *Decoder {
	scan := scanner.Scanner{
		Mode: scanner.GoTokens,
	}
	scan.Init(r)
	return &Decoder{
		&lexer{
			scan: scan,
		},
	}
}

func (d *Decoder) Decode(out interface{}) (err error) {
	d.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error occured, %v: %v", d.lex.scan.Position, x)
		}
	}()
	read(d.lex, reflect.ValueOf(out).Elem())
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{
		scan: scanner.Scanner{
			Mode: scanner.GoTokens,
		},
	}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error occured, %v: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func encode(buf *bytes.Buffer, v reflect.Value, indent int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		encode(buf, v.Elem(), indent)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		indent++
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte('\n')
				buf.WriteString(getIndentSpaces(indent))
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			zeroInterface := reflect.Zero(v.Field(i).Type()).Interface()
			if reflect.DeepEqual(v.Field(i).Interface(), zeroInterface) {
				continue
			}
			if i > 0 {
				buf.WriteByte('\n')
				buf.WriteString(getIndentSpaces(indent))
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent+1); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		indent++
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte('\n')
				buf.WriteString(getIndentSpaces(indent))
			}
			buf.WriteByte('(')
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	case reflect.Bool: // t or nil
		if v.Bool() {
			buf.WriteByte('t')
		} else {
			buf.WriteString("nil")
		}

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%g", v.Float())

	case reflect.Complex64, reflect.Complex128: // complex
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%g %g)", real(c), imag(c))

	case reflect.Interface:
		buf.WriteByte('(')
		indent++
		fmt.Fprintf(buf, "(%q ", reflect.Indirect(v).Type())
		indent++
		buf.WriteByte(' ')
		indent++
		encode(buf, reflect.Indirect(v).Elem(), indent)
		buf.WriteByte(')')

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func getIndentSpaces(indent int) string {
	return strings.Repeat(" ", indent*4)
}
