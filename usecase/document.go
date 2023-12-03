package usecase

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
)

type Document struct {
	docPort port.Document
}

func NewDocument(i *do.Injector) (*Document, error) {
	docPort := do.MustInvoke[port.Document](i)
	return &Document{
		docPort: docPort,
	}, nil
}

func (d *Document) FindContent(id edinet.DocumentId, fileType edinet.FileType) (edinet.DocumentFile, error) {
	document, err := d.docPort.Get(id, fileType)
	if err != nil {
		return edinet.DocumentFile{}, err
	}
	return document, nil
}

func (d *Document) FindAllDocs() ([]edinet.DocumentFile, error) {
	panic("todo")
}
