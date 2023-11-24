package edinet

import (
	"fmt"
	"github.com/tkitsunai/edinet-go/core"
)

type RequestType int

func (r RequestType) String() string {
	return fmt.Sprintf("%d", r)
}

const (
	MetaDataOnly RequestType = iota + 1
	MetaDataAndDocuments
)

type DocumentListRequestParameter struct {
	FileDate core.FileDate
	Type     RequestType
}

// must be generated struct "DocumentListResponse"
// gen_document_list.go
type DocumentContentResponses []*DocumentListResponse
type DocumentListResponses struct {
	List []DocumentListResponse `json:"items"`
}
