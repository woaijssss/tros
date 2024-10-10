package http

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

// ResToObj Structuring the JSON response of HTTP into a struct
func ResToObj(resp *Response, target any) error {
	bs, err := readResponse(resp)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return json.Unmarshal(bs, target)
}

// ReqToObj Structuring the JSON request of HTTP into a struct
func ReqToObj(req *Request, target any) error {
	bs, err := readRequest(req)
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
	bs, err := readResponse(resp)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return xml.Unmarshal(bs, target)
}

// ReqXmlToObj Structuring HTTP XML requests into structures
func ReqXmlToObj(req *Request, target any) error {
	bs, err := readRequest(req)
	if err != nil {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return xml.Unmarshal(bs, target)
}

// Read byte stream data from HTTP requests
func readRequest(req *Request) ([]byte, error) {
	if req == nil {
		return nil, nil
	}
	if req.Body == nil {
		return nil, nil
	}
	defer req.Body.Close()
	return io.ReadAll(req.Body)
}

// Read byte stream data of HTTP response
func readResponse(resp *Response) ([]byte, error) {
	if resp == nil {
		return nil, nil
	}
	if resp.Body == nil {
		return nil, nil
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
