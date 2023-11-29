package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/conf"
	"github.com/tkitsunai/edinet-go/datastore"
	"github.com/tkitsunai/edinet-go/usecase"
	"net"
)

type Server struct {
	app         *fiber.App
	storeEngine datastore.Engine
	i           *do.Injector
}

func NewServer(storeEngine datastore.Engine, injector *do.Injector) *Server {
	s := &Server{
		app: fiber.New(fiber.Config{
			Prefork:      false,
			AppName:      "EDINET-GO",
			ServerHeader: "edinet-go",
		}),
		storeEngine: storeEngine,
		i:           injector,
	}
	s.setHandlers()
	return s
}

func (s *Server) setHandlers() {
	overview := do.MustInvoke[*usecase.Overview](s.i)
	document := do.MustInvoke[*usecase.Document](s.i)
	docResources := NewDocumentsResource(overview, document)

	s.app.Get("/documents", docResources.GetDocumentsByTerm)
	s.app.Get("/documents/:id", docResources.GetDocument)
}

func (s *Server) Run() error {
	return s.app.Listen(net.JoinHostPort("", conf.LoadServerConfig().Port))
}
