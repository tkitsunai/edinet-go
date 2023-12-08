package usecase

import (
	"fmt"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
)

type PassThrough struct {
	overview port.Overview
	document port.Document
}

func NewPassThrough(i *do.Injector) (*PassThrough, error) {
	return &PassThrough{
		overview: do.MustInvoke[port.Overview](i),
		document: do.MustInvoke[port.Document](i),
	}, nil
}

func (p *PassThrough) DocumentMetas(
	date core.Date,
	typ edinet.RequestType,
) (edinet.EdinetDocumentResponse, error) {
	switch typ {
	case edinet.MetaDataAndDocuments, edinet.MetaDataOnly:
		return p.overview.GetRaw(date, typ)
	default:
		return edinet.EdinetDocumentResponse{}, fmt.Errorf("request type not match")
	}
}
