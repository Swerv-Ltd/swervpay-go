package swervpay

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func GenerateURLPath(path string, query interface{}) string {
	v := url.Values{}

	q := reflect.ValueOf(query)
	if q.Kind() == reflect.Ptr {
		q = q.Elem() // Dereference the pointer to get the underlying struct
	}

	for i := 0; i < q.NumField(); i++ {
		field := q.Type().Field(i)
		value := q.Field(i).Interface()
		v.Add(strings.ToLower(field.Name), strconv.Itoa(value.(int)))
	}

	urlQuery := v.Encode()

	// Always use ? to start query parameters, unless the path already contains a ?
	if strings.Contains(path, "?") {
		path += "&" + urlQuery
	} else {
		path += "?" + urlQuery
	}

	return path
}
