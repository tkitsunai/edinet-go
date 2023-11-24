package core

type ContentType string

func (c ContentType) String() string {
	return string(c)
}

const (
	ZIPContentType  ContentType = "application/octet-stream"
	PDFContentType  ContentType = "application/pdf"
	JSONContentType ContentType = "application/json; charset=utf-8"
)
