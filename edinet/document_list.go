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
	FileDate core.FileDate
	Type     RequestType
}

type File struct {
	Name       string
	Extension  string
	DocumentId DocumentId
	Content    []byte
}

func (f *File) NameWithExtension() string {
	return f.Name + f.Extension
}

// DocumentContentResponses
// must be generated struct "DocumentListResponse"
// gen_document_list.go
type DocumentContentResponses []*DocumentListResponse
type DocumentListResponses struct {
	List []DocumentListResponse `json:"items"`
}
