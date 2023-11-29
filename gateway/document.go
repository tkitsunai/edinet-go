package gateway

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
)

type Document struct {
	c *edinet.Client
}

func NewDocument(i *do.Injector) (port.Document, error) {
	client := do.MustInvoke[*edinet.Client](i)
	return &Document{c: client}, nil
}

func (d *Document) Get(id edinet.DocumentId, fileType edinet.FileType) (edinet.File, error) {
	return d.c.RequestDocument(id, fileType)
}
