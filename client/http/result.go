package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

// ResToObj Structuring the JSON response of HTTP into a struct
func ResToObj(resp *Response, target any) error {
	bs, err := read(resp)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return json.Unmarshal(bs, target)
}

// ResXmlToObj Structuring HTTP XML responses into structures
func ResXmlToObj(resp *Response, target any) error {
	bs, err := read(resp)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return xml.Unmarshal(bs, target)
}

// 读取http响应的字节流数据
func read(resp *Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
