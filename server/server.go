package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/edinet-go/conf"
)

type Server struct {
	app *fiber.App
}

func init() {
}

func (s *Server) setHandlers() {
	docResources := NewDocumentsResource(conf.LoadConfig().ApiKey)
	s.app.Get("/documents", docResources.GetDocumentsByTerm)
}

func (s *Server) Run() error {
	return s.app.Listen(":" + conf.LoadServerConfig().Port)
}

func NewServer() *Server {
	s := &Server{
		app: fiber.New(),
	}
	s.setHandlers()
	return s
}
