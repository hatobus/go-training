package ex12_5

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
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
		indent++
		buf.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(",\n")
				buf.WriteString(getIndentSpaces(indent))
			}
			if err := encode(buf, v.Index(i), indent); err != nil {
				return err
			}
		}
		buf.WriteString("]")

	case reflect.Struct: // ((name value) ...)
		indent++
		buf.WriteString("{\n")
		buf.WriteString(getIndentSpaces(indent))
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteString(",\n")
				buf.WriteString(getIndentSpaces(indent))
			}
			fmt.Fprintf(buf, "%s: ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), indent); err != nil {
				return err
			}
		}
		buf.WriteString("\n}")

	case reflect.Map: // ((key value) ...)
		indent++
		buf.WriteString("{\n")
		buf.WriteString(getIndentSpaces(indent))
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteString(",\n")
				buf.WriteString(getIndentSpaces(indent))
			}
			if err := encode(buf, key, indent); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), indent); err != nil {
				return err
			}
		}
		buf.WriteString("\n}")

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

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func getIndentSpaces(indent int) string {
	return strings.Repeat(" ", indent*4)
}
