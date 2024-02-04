package http

import (
	"encoding/json"
	"io"
)

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
