package usecase

import (
	"fmt"
	"proxy-example/domain"
)

// Client coordinates request sending.
type Client struct {
	logger domain.Logger
}

// NewClient builds a client use case with the provided logger.
func NewClient(logger domain.Logger) *Client {
	return &Client{logger: logger}
}

// SendRequests issues repeated requests via the provided server.
func (c *Client) SendRequests(server domain.Server, url string, count int) {
	for i := 0; i < count; i++ {
		code, body := server.HandleRequest(url, "GET")
		c.logger.Log(fmt.Sprintf("Request %d: Code=%d, Body=%s", i+1, code, body))
	}
}
