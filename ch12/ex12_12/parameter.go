package ex12_12

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var (
	phoneNumberValidate = regexp.MustCompile(`^0[789]0-\d{4}-\d{4}$`)             // 000-1111-2222
	emailValidate       = regexp.MustCompile(`^[a-zA-Z0-9-_.]+@[a-zA-Z0-9-_.]+$`) // example@ex.com
	cardVISAValidate    = regexp.MustCompile(`4[0-9]{12}(?:[0-9]{3})?`)           // VISA card
)

// Upk populates the fields of the struct pointed to by ptr
// from the HTTP request parameters in req.
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	fields := make(map[string]reflect.Value)
	validateTags := make(map[string]string)
	v := reflect.ValueOf(ptr).Elem() // the struct variable
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		validateTarget := tag.Get("validation")
		if validateTarget != "" {
			validateTags[name] = validateTarget
			//fmt.Println(validateTarget)
			//if err := validate(validateTarget, v.Field(i).Interface()); err != nil {
			//	return err
			//}
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
		if tag, ok := validateTags[name]; ok {
			if err := validate(tag, f.String()); err != nil {
				return err
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

func validate(target string, elem interface{}) error {
	v, _ := elem.(string)
	switch target {
	case "phone":
		if !phoneNumberValidate.MatchString(v) {
			return errors.New("invalid phone number")
		}
	case "card":
		if !cardVISAValidate.MatchString(v) {
			return errors.New("invalid card number")
		}
	case "email":
		if !emailValidate.MatchString(v) {
			return errors.New("invalid email format")
		}
	}
	return nil
}
