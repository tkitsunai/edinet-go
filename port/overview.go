package port

import (
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
)

type Overview interface {
	Get(date core.FileDate) (*edinet.DocumentListResponse, error)
	GetByStore(date core.FileDate) (*edinet.DocumentListResponse, error)
}
