package swervpay

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func GenerateURLPath(path string, query interface{}) string {
	if query == nil {
		return path
	}

	v := url.Values{}

	q := reflect.ValueOf(query)
	for q.Kind() == reflect.Ptr {
		if q.IsNil() {
			return path
		}
		q = q.Elem() // Dereference the pointer to get the underlying struct
	}
	if q.Kind() != reflect.Struct {
		return path
	}

	for i := 0; i < q.NumField(); i++ {
		field := q.Type().Field(i)
		if field.PkgPath != "" {
			continue
		}

		name := strings.ToLower(field.Name)
		if tag := strings.Split(field.Tag.Get("json"), ",")[0]; tag != "" && tag != "-" {
			name = tag
		}

		value := q.Field(i).Interface()
		switch value := value.(type) {
		case int:
			v.Add(name, strconv.Itoa(value))
		case string:
			if value != "" {
				v.Add(name, value)
			}
		}
	}

	urlQuery := v.Encode()
	if urlQuery == "" {
		return path
	}

	if strings.Contains(path, "?") {
		path += "&" + urlQuery
	} else {
		path += "?" + urlQuery
	}

	return path
}
