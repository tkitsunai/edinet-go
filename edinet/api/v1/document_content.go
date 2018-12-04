package v1

import (
	"errors"
	"fmt"

	"github.com/tkitsunai/edinet-go/edinet/core"
)

type DocumentRequestType int

const (
	AuditReport DocumentRequestType = iota + 1
	PDFRequestType
	Attachment
	EnglishDocument
)

type DocumentRequestParameter struct {
	Type DocumentRequestType
}

type DocumentResponse struct {
	BinaryData []byte
	Header     core.ContentType
}

func FactoryByDocumentType(typ DocumentRequestType) (*DocumentResponse, error) {
	switch typ {
	case AuditReport:
		return &DocumentResponse{
			BinaryData: nil, Header: core.ZIPContentType,
		}, nil
	case PDFRequestType:
		return &DocumentResponse{
			BinaryData: nil, Header: core.PDFContentType,
		}, nil
	case Attachment:
		return &DocumentResponse{
			BinaryData: nil, Header: core.ZIPContentType,
		}, nil
	case EnglishDocument:
		return &DocumentResponse{
			BinaryData: nil, Header: core.ZIPContentType,
		}, nil
	}
	return &DocumentResponse{
		BinaryData: nil, Header: core.JSONContentType,
	}, errors.New(fmt.Sprintf("No such document type. %v", typ))
}
