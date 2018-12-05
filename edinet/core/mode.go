package core

type Mode string

const (
	Release Mode = "release"
	Debug   Mode = "debug"
)

type V1Engine interface {
	RequestDocumentContent() error
	RequestDocumentList() error
	SetMode(Mode) (V1Engine, error)
}
