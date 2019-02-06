package v1

type Engine interface {
	RequestDocumentContent(RequestType) error
	RequestDocumentList(RequestType) error
	RequestDocumentListByParameter(DocumentListRequestParameter) error
}
