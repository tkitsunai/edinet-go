package port

import (
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
)

type Overview interface {
	Get(date core.Date) (*edinet.DocumentListResponse, error)
	GetByStore(date core.Date) (*edinet.DocumentListResponse, error)
}
