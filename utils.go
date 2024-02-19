package swervpay

import (
	"net/url"
	"reflect"
	"strconv"
)

func GenerateURLPath(path string, query interface{}) string {
	v := url.Values{}

	q := reflect.ValueOf(query)
	for i := 0; i < q.NumField(); i++ {
		field := q.Type().Field(i)
		value := q.Field(i).Interface()
		v.Add(field.Name, strconv.Itoa(value.(int)))
	}

	urlQuery := v.Encode()

	if url.PathEscape(path) != path {
		path += "&" + urlQuery
	} else {
		path += "?" + urlQuery
	}

	return path
}
