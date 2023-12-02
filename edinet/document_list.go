package edinet

import (
	"fmt"
	"github.com/tkitsunai/edinet-go/core"
	"strings"
)

type RequestType int
type FileType int

func (r RequestType) String() string {
	return fmt.Sprintf("%d", r)
}

func (f FileType) String() string {
	return fmt.Sprintf("%d", f)
}

type DocumentId string

func (f DocumentId) String() string {
	return string(f)
}

const (
	MetaDataOnly RequestType = iota + 1
	MetaDataAndDocuments
)

const (
	Unknown FileType = iota
	XBRL
	PDF
	AlternativeDoc
	EnglishDoc
	CSV
)

func NewFileTypeByName(name string) FileType {
	if len(name) == 0 {
		return Unknown
	}
	switch strings.ToUpper(name) {
	case "XBRL":
		return XBRL
	case "PDF":
		return PDF
	case "CSV":
		return CSV
	}
	return Unknown
}

type DocumentListRequestParameter struct {
	FileDate core.Date
	Type     RequestType
}

type DocumentFile struct {
	Name       string
	Extension  string
	DocumentId DocumentId
	Content    []byte
}

func (f *DocumentFile) NameWithExtension() string {
	return f.Name + f.Extension
}

// EdinetResponses
// must be generated struct "EdinetDocumentResponse"
// gen_edinet_response.go
type EdinetResponses struct {
	Items []EdinetDocumentResponse `json:"items"`
}
