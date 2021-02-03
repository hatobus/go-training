package ex12_11

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Pack(ptr interface{}) (url.URL, error) {
	v := reflect.ValueOf(ptr).Elem()
	if v.Type().Kind() != reflect.Struct {
		return url.URL{}, fmt.Errorf("%v is not a struct", ptr)
	}
	values := &url.Values{}
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag
		tagHttp := tag.Get("http")
		if tagHttp == "" {
			tagHttp = strings.ToLower(field.Name)
		}
		values.Add(tagHttp, fmt.Sprintf("%v", v.Field(i)))
	}
	return url.URL{RawQuery: values.Encode()}, nil
}
