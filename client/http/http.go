package http

import (
	"bytes"
	"context"
	"encoding/xml"
	trlogger "github.com/woaijssss/tros/logx"
	"io"
	"net/http"
)

func NewHttpClient() *Client {
	return httpPool.getConn()
}

func (c *Client) SetHeader(k, v string) {
	c.header[k] = v
}

func (c *Client) Post(ctx context.Context, url string, body io.Reader) (*Response, error) {
	return c.send(ctx, http.MethodPost, url, body)
}

func (c *Client) PostXml(ctx context.Context, url string, data any) (*Response, error) {
	xmlData, err := xml.Marshal(data)
	if err != nil {
		return nil, err
	}
	return c.sendXml(ctx, http.MethodPost, url, xmlData)
}

func (c *Client) Get(ctx context.Context, url string) (*Response, error) {
	return c.send(ctx, http.MethodGet, url, nil)
}

func (c *Client) GetWithReader(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.send(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return readResponse(resp)
}

func (c *Client) GetHeader(k string) string {
	v, ok := c.header[k]
	if !ok {
		return ""
	}
	return v
}

// json请求
func (c *Client) send(ctx context.Context, method, url string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		trlogger.Errorf(ctx, "new http json request with %s err: [%+v]", method, err)
		return nil, err
	}

	for k, v := range c.header {
		req.Header.Set(k, v)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		trlogger.Errorf(ctx, "do http json with %s err: [%+v]", method, err)
		return nil, err
	}

	return &Response{resp}, nil
}

// xml请求
func (c *Client) sendXml(ctx context.Context, method, url string, xmlData []byte) (*Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(xmlData))
	if err != nil {
		trlogger.Errorf(ctx, "new http xml request with %s err: [%+v]", method, err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/xml")

	resp, err := c.hc.Do(req)
	if err != nil {
		trlogger.Errorf(ctx, "do http xml with %s err: [%+v]", method, err)
		return nil, err
	}

	return &Response{resp}, nil
}

type Client struct {
	hc     *http.Client
	header map[string]string
}

type Request struct {
	*http.Request
}

type Response struct {
	*http.Response
	//Json map[string]interface{} // 返回值的json格式
}
