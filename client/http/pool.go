package http

import (
	"net/http"
	"sync"
)

var httpPool = newHttpPool()

type pool struct {
	pool   sync.Pool
	client *http.Client
}

func newHttpPool() *pool {
	return &pool{
		pool: sync.Pool{
			New: func() interface{} {
				return &http.Client{}
			},
		},
	}
}

func (p *pool) getConn() *Client {
	return &Client{
		hc:     httpPool.pool.Get().(*http.Client),
		header: make(map[string]string, 3),
	}
}

func (p *pool) putConn(c *http.Client) {
	httpPool.pool.Put(c)
}

func (p *pool) do(req *http.Request) {

}
