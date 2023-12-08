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
	ALL
)

func AllFileType() []FileType {
	return []FileType{
		XBRL, PDF, AlternativeDoc, EnglishDoc, CSV,
	}
}

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
	case "ALL":
		return ALL
	}
	return Unknown
}

type DocumentListRequestParameter struct {
	FileDate core.Date
	Type     RequestType
}

type DocumentFile struct {
	Name        string
	Extension   string
	DocumentId  string
	ContentType string
	Content     []byte
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
