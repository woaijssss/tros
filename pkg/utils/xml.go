package utils

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// StructToXMLKeyValueSorted
/*
	Concatenate the structure tag into a key=value format for parameters, and sort them in ASCII lexicographic order by parameter name
*/
func StructToXMLKeyValueSorted(s interface{}) string {
	v := reflect.ValueOf(s)
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return ""
	}

	parts := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		xmlTag := t.Field(i).Tag.Get("xml")
		if xmlTag != "" {
			value := fmt.Sprintf("%v", fieldValue)
			parts = append(parts, fmt.Sprintf("%s=%s", xmlTag, value))
		}
	}

	sort.Strings(parts)
	return strings.Join(parts, "&")
}
