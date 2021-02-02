package ex12_2

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
)

func Display(name string, x interface{}) string {
	return display(name, reflect.ValueOf(x), 0)
}

func display(path string, v reflect.Value, step int) string {
	if step > 3 {
		return fmt.Sprintf("%s = %s\n", path, formatAtom(v))
	}
	buf := &bytes.Buffer{}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(buf, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			fmt.Fprint(buf, display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), step+1))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			fmt.Fprint(buf, display(fieldPath, v.Field(i), step+1))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			fmt.Fprint(buf,
				display(
					fmt.Sprintf("%s[%s]", path, resolveMapKey(key)),
					v.MapIndex(key),
					step+1,
				),
			)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Fprint(buf, display(fmt.Sprintf("(*%s)", path), v.Elem(), step+1))
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(buf, "%s = nil\n", path)
		} else {
			fmt.Fprintf(buf, "%s.type = %s\n", path, v.Elem().Type())
			fmt.Fprint(buf, display(path+".value", v.Elem(), step+1))
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
