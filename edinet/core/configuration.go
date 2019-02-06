package core

import (
	"net/url"
	"os"
)

type Configuration interface {
	RequestBaseUri() url.URL
}

type EdinetConfig struct {
	edinetApiDocumentsUrl       string
	edinetApiDocumentContentUrl string
}

const (
	EDINET_API_DOCUMENTS_BASE_URL        = "EDINET_API_DOCUMENTS_BASE_URL"
	EDINET_API_DOCUMENT_CONTENT_BASE_URL = "EDINET_API_DOCUMENT_CONTENT_BASE_URL"
)

// constructor setting by environment variables
func NewEdinetConfig() *EdinetConfig {
	return &EdinetConfig{
		edinetApiDocumentsUrl:       os.Getenv(EDINET_API_DOCUMENTS_BASE_URL),
		edinetApiDocumentContentUrl: os.Getenv(EDINET_API_DOCUMENT_CONTENT_BASE_URL),
	}
}

func (e *EdinetConfig) RequestBaseUri() url.URL {
	panic("not implement")
	return url.URL{}
}
