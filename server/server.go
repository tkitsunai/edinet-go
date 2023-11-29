package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/edinet-go/conf"
	"github.com/tkitsunai/edinet-go/datastore"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/usecase"
	"net"
)

type Server struct {
	app         *fiber.App
	storeEngine datastore.Engine
}

func NewServer(storeEngine datastore.Engine) *Server {
	s := &Server{
		app: fiber.New(fiber.Config{
			Prefork:      false,
			AppName:      "EDINET-GO",
			ServerHeader: "edinet-go",
		}),
		storeEngine: storeEngine,
	}
	s.setHandlers()
	return s
}

func (s *Server) setHandlers() {
	key := conf.LoadConfig().ApiKey
	client := edinet.NewClient(key)

	overview := usecase.NewOverview(client)
	document := usecase.NewDocument(client)

	docResources := NewDocumentsResource(overview, document)

	s.app.Get("/documents", docResources.GetDocumentsByTerm)
	s.app.Get("/documents/:id", docResources.GetDocument)
}

func (s *Server) Run() error {
	return s.app.Listen(net.JoinHostPort("", conf.LoadServerConfig().Port))
}
