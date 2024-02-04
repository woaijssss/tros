package utils

import (
	"encoding/json"
	"fmt"
	"os"
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

func StringToJson[T any](s string) (T, error) {
	var v T
	err := json.Unmarshal([]byte(s), &v)
	return v, err
}

func ByteToJson[T any](b []byte) (T, error) {
	var v T
	err := json.Unmarshal(b, &v)
	return v, err
}
