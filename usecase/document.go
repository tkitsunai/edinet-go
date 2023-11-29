package usecase

import "github.com/tkitsunai/edinet-go/edinet"

type Document struct {
	Client *edinet.Client
}

func NewDocument(client *edinet.Client) *Document {
	return &Document{
		Client: client,
	}
}

func (c *Document) FindContentById(id edinet.DocumentId, fileType edinet.FileType) (edinet.File, error) {
	document, err := c.Client.RequestDocument(id, fileType)
	if err != nil {
		return edinet.File{}, err
	}
	return document, nil
}
