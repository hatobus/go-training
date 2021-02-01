package ex12_1

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(path string, v reflect.Value) string {
	buf := &bytes.Buffer{}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(buf, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			fmt.Fprint(buf, Display(fmt.Sprintf("%s[%d]", path, i), v.Index(i)))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			fmt.Fprint(buf, Display(fieldPath, v.Field(i)))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fmt.Fprint(buf,
				Display(
					fmt.Sprintf("%s[%s]", path, resolveMapKey(key)),
					v.MapIndex(key),
				),
			)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Fprint(buf, Display(fmt.Sprintf("(*%s)", path), v.Elem()))
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(buf, "%s = nil\n", path)
		} else {
			fmt.Fprintf(buf, "%s.type = %s\n", path, v.Elem().Type())
			fmt.Fprint(buf, Display(path+".value", v.Elem()))
		}
	default:
		fmt.Fprintf(buf, "%s = %s\n", path, formatAtom(v))
	}
	return buf.String()
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

func resolveMapKey(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Struct:
		buf := &bytes.Buffer{}
		buf.WriteString("{")

		fields := v.NumField()
		for i := 0; i < fields; i++ {
			fmt.Fprintf(buf, "%s: %s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
			if i != fields-1 {
				buf.WriteString(", ")
			}
		}

		buf.WriteString("}")
		return buf.String()
	case reflect.Array, reflect.Slice:
		buf := &bytes.Buffer{}
		buf.WriteString("[")

		length := v.Len()
		for i := 0; i < length; i++ {
			buf.WriteString(formatAtom(v.Index(i)))

			if i != length-1 {
				buf.WriteString(", ")
			}
		}

		buf.WriteString("]")
		return buf.String()
	default:
		return formatAtom(v)
	}
}
