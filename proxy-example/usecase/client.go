package usecase

import (
	"fmt"
	"proxy-example/domain"
)

type Client struct {
	logger domain.Logger
}

func NewClient(logger domain.Logger) *Client {
	return &Client{logger: logger}
}

func (c *Client) SendRequests(server domain.Server, url string, count int) {
	for i := 0; i < count; i++ {
		code, body := server.HandleRequest(url, "GET")
		c.logger.Log(fmt.Sprintf("Request %d: Code=%d, Body=%s", i+1, code, body))
	}
}
