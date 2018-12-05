package edinet

import (
	"github.com/tkitsunai/edinet-go/edinet/core"
)

func defaultClient() *client {
	return &client{
		mode: core.Release,
	}
}

func NewDefault() core.V1Engine {
	return defaultClient()
}

type client struct {
	mode core.Mode
}

func (c *client) SetMode(mode core.Mode) (core.V1Engine, error) {
	c.mode = mode
	return c, nil
}

func (c *client) RequestDocumentList() error {
	return nil
}

func (c *client) RequestDocumentContent() error {
	return nil
}
