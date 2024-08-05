package http

import (
	"context"
	trlogger "gitee.com/idigpower/tros/logx"
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

func (c *Client) Get(ctx context.Context, url string) (*Response, error) {
	return c.send(ctx, http.MethodGet, url, nil)
}

func (c *Client) GetWithReader(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.send(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	return read(resp)
}

func (c *Client) GetHeader(k string) string {
	v, ok := c.header[k]
	if !ok {
		return ""
	}
	return v
}

func (c *Client) send(ctx context.Context, method, url string, body io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		trlogger.Errorf(ctx, "new http request with %s err: [%+v]", method, err)
		return nil, err
	}

	for k, v := range c.header {
		req.Header.Set(k, v)
	}

	resp, err := c.hc.Do(req)
	if err != nil {
		trlogger.Errorf(ctx, "do http with %s err: [%+v]", method, err)
		return nil, err
	}

	return &Response{resp}, nil
}

type Client struct {
	hc     *http.Client
	header map[string]string
}

type Response struct {
	*http.Response
	//Json map[string]interface{} // 返回值的json格式
}
