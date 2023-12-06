package port

import (
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
)

type Document interface {
	Get(id core.DocumentId, fileType edinet.FileType) (edinet.DocumentFile, error)
}
