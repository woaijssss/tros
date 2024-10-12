package utils

import (
	"encoding/json"
	"fmt"
	"github.com/woaijssss/tros/trerror"
	"os"
	"reflect"
	"sort"
	"strings"
)

func SaveJson(marshal []byte, jsonFile string) (err error) {
	create, err := os.Create(jsonFile)
	if err != nil {
		fmt.Println("cretre error", err)
		return
	}
	// 用后关闭
	defer create.Close()

	_, err = create.Write(marshal)
	if err != nil {
		fmt.Println("write error", err)
		return
	}

	return
}

func ToJsonByte(v any) ([]byte, error) {
	return json.Marshal(v)
}

func ToJsonString(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// ToJsonSortedString Sort the structure into a JSON string in ascending order according to the first letter ASCII code
func ToJsonSortedString(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	var tempMap map[string]interface{}
	err = json.Unmarshal(b, &tempMap)
	if err != nil {
		return "", err
	}

	keys := make([]string, 0, len(tempMap))
	for k := range tempMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	sortedMap := make(map[string]interface{})
	for _, k := range keys {
		sortedMap[k] = tempMap[k]
	}

	jsonBytes, err := json.Marshal(sortedMap)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(jsonBytes), nil
}

func StringToJson[T any](s string) (T, error) {
	var v T
	if s == "" {
		return v, nil
	}
	err := json.Unmarshal([]byte(s), &v)
	return v, err
}

func ByteToJson[T any](b []byte) (T, error) {
	var v T
	err := json.Unmarshal(b, &v)
	return v, err
}

func MapToJson[T any](m map[string]interface{}) (T, error) {
	var t T
	b, err := json.Marshal(m)
	if err != nil {
		return t, trerror.DefaultTrError(fmt.Sprintf("marshal map err: [%+v]", err))
	}
	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, trerror.DefaultTrError(fmt.Sprintf("unmarshal to struct err: [%+v]", err))
	}
	return t, nil
}

// StructToKeyValueSorted
/*
	Concatenate the structure into a key=value format for parameters, and sort them in ASCII lexicographic order by parameter name
*/
func StructToKeyValueSorted(s any) string {
	v := reflect.ValueOf(s)
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return ""
	}

	parts := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		key := t.Field(i).Tag
		value := fmt.Sprintf("%v", fieldValue)
		parts = append(parts, fmt.Sprintf("%s=%s", key, value))
	}

	sort.Strings(parts)
	return strings.Join(parts, "&")
}

// StructToJSONKeyValueSorted
/*
	Concatenate the structure tag into a key=value format for parameters, and sort them in ASCII lexicographic order by parameter name
*/
func StructToJSONKeyValueSorted(s interface{}) string {
	v := reflect.ValueOf(s)
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return ""
	}

	parts := make([]string, 0)
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i).Interface()
		jsonTag := t.Field(i).Tag.Get("json")
		if jsonTag != "" {
			value := fmt.Sprintf("%v", fieldValue)
			parts = append(parts, fmt.Sprintf("%s=%s", jsonTag, value))
		}
	}

	sort.Strings(parts)
	return strings.Join(parts, "&")
}
