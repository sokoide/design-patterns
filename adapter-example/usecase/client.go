package usecase

import (
	"adapter-example/domain"
)

// Client coordinates adapter usage for the use case.
type Client struct {
	logger domain.Logger
}

// NewClient builds a client use case with the provided logger.
func NewClient(logger domain.Logger) *Client {
	return &Client{logger: logger}
}

// InsertLightningConnectorIntoComputer uses a computer via the common interface.
func (c *Client) InsertLightningConnectorIntoComputer(computer domain.Computer) {
	c.logger.Log("Client inserts Lightning connector into computer.")
	computer.InsertIntoLightningPort()
}
