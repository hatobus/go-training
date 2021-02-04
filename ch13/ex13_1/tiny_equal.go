package ex13_1

import (
	"reflect"
	"unsafe"
)

var mul = 1e9

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return tinyEqual(float64(x.Int()), float64(y.Int()))
	case reflect.Float32, reflect.Float64:
		return tinyEqual(x.Float(), y.Float())
	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true
	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, key := range x.MapKeys() {
			if equal(x.MapIndex(key), y.MapIndex(key), seen) {
				return false
			}
		}
		return true
	}
	panic("unexpected")
}

func tinyEqual(x, y float64) bool {
	if x == y {
		return true
	}
	var diff float64
	if x > y {
		diff = x - y
	} else {
		diff = y - x
	}
	d := diff * mul
	if d < x && d < y {
		return true
	}
	return false
}
