package core

type Mode string

const (
	Release Mode = "release"
	Debug   Mode = "debug"
)

type Engine interface {
	RequestDocumentContent() error
	RequestDocumentList() error
	SetMode(Mode) (Engine, error)
}
