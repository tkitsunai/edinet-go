package v1

import "errors"

type RequestType int

const (
	MetaDataOnly         RequestType = iota +1
	MetaDataAndDocuments
)

type DocumentListRequestParameter struct {
	FileDate FileDate
	Type     RequestType
}

type FileDate string

func (f FileDate) Validate() error {
	if len(f) == 0 {
		return errors.New("RequiredParameter")
	}
	return nil
}

// must be generated struct "DocumentListResponse"
type DocumentContentResponses []*DocumentListResponse
