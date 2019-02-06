package edinet

import (
	"github.com/tkitsunai/edinet-go/edinet/api/v1"
	"github.com/tkitsunai/edinet-go/edinet/core"
)

type Mode string

const (
	Release Mode = "release"
	Debug   Mode = "debug"
)

func defaultClient() *client {
	return &client{
		mode: Release,
	}
}

func NewDefault() v1.Engine {
	return defaultClient()
}

type client struct {
	mode Mode

	requestForDocumentList    core.Configuration
	requestForDocumentContent core.Configuration
}

func (c *client) RequestDocumentListByParameter(parameter v1.DocumentListRequestParameter) error {
	err := parameter.FileDate.Validate()

	if err != nil {
		return err
	}

	//core.NewHttpClient().ExecuteGetJsonType(core.NewEdinetConfig().RequestBaseUri().String())
	panic("implement me")
}

func (c *client) RequestDocumentList(requestType v1.RequestType) error {
	panic("implement me")
}

func (c *client) RequestDocumentContent(requestType v1.RequestType) error {
	panic("implement me")
}
