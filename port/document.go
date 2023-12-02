package port

import "github.com/tkitsunai/edinet-go/edinet"

type Document interface {
	Get(id edinet.DocumentId, fileType edinet.FileType) (edinet.DocumentFile, error)
}
