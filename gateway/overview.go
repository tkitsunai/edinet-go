package gateway

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
)

type Overview struct {
	c *edinet.Client
}

func NewOverview(i *do.Injector) (port.Overview, error) {
	c := do.MustInvoke[*edinet.Client](i)
	return &Overview{c: c}, nil
}

func (o *Overview) Get(date core.FileDate) (*edinet.DocumentListResponse, error) {
	return o.c.RequestDocumentList(date)
}
