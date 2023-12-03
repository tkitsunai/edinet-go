package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/conf"
	"github.com/tkitsunai/edinet-go/datastore"
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
	s.app.Use(pprof.New())
	s.setHandlers()
	return s
}

func (s *Server) setHandlers() {
	docResources := do.MustInvoke[*Documents](s.i)
	companyResoures := do.MustInvoke[*Company](s.i)

	s.app.Get("/documents", docResources.GetDocumentsByTerm)
	s.app.Post("/documents", docResources.StoreDocumentsByTerm)
	s.app.Get("/documents/:id", docResources.GetDocument)
	s.app.Get("/companies", companyResoures.FindCompanies)
	s.app.Get("/companies/:id", companyResoures.FindCompany)
}

func (s *Server) Run() error {
	return s.app.Listen(net.JoinHostPort("", conf.LoadServerConfig().Port))
}
