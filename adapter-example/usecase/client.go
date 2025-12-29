package usecase

import (
	"adapter-example/domain"
)

type Client struct {
	logger domain.Logger
}

func NewClient(logger domain.Logger) *Client {
	return &Client{logger: logger}
}

func (c *Client) InsertLightningConnectorIntoComputer(com domain.Computer) {
	c.logger.Log("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
