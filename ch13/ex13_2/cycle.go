package ex13_2

import (
	"reflect"
	"unsafe"
)

type ptr struct {
	x unsafe.Pointer
	t reflect.Type
}

func Cycle(x interface{}) bool {
	seen := make(map[ptr]bool)
	return cycle(reflect.ValueOf(x), seen)
}

func cycle(x reflect.Value, seen map[ptr]bool) bool {
	if x.CanAddr() {
		p := ptr{
			unsafe.Pointer(x.UnsafeAddr()),
			x.Type(),
		}
		if seen[p] {
			return true
		}
		seen[p] = true
	}
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return cycle(x.Elem(), seen)
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cycle(x.Index(i), seen) {
				return true
			}
		}
		return false
	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if cycle(x.Field(i), seen) {
				return true
			}
		}
		return false
	case reflect.Map:
		for _, key := range x.MapKeys() {
			if cycle(x.MapIndex(key), seen) || cycle(key, seen) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
