package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/conf"
	myLogger "github.com/tkitsunai/edinet-go/logger"
	"net"
	"strings"
)

type Server struct {
	app *fiber.App
	i   *do.Injector
	cfg Config
}

type ServerMode string

const (
	DEVELOPMENT_MODE = ServerMode("DEV")
	PRODUCTION_MODE  = ServerMode("PRODUCION")
)

func OfMode(mode string) ServerMode {
	switch strings.ToUpper(mode) {
	case "DEVELOPMENT", "DEV":
		return DEVELOPMENT_MODE
	case "PRODUCTION", "PROD", "PRD":
		return PRODUCTION_MODE
	default:
		return DEVELOPMENT_MODE
	}
}

type Config struct {
	Mode ServerMode
}

func NewServer(injector *do.Injector, cfg Config) *Server {
	s := &Server{
		app: fiber.New(fiber.Config{
			Prefork:      false,
			AppName:      "EDINET-GO",
			ServerHeader: "edinet-go",
		}),
		i: injector,
	}

	s.cfg = cfg
	s.app.Use(myLogger.RequestLogging())
	s.setHandlers()
	return s
}

func (s *Server) setHandlers() {
	docResources := do.MustInvoke[*Overview](s.i)
	companyResources := do.MustInvoke[*Company](s.i)
	edinetResources := do.MustInvoke[*EdinetRaw](s.i)
	s.app.Get("/_raw/api/v2/documents.json", edinetResources.GetMetaDataByDate)
	s.app.Get("/_raw/api/v2/documents/:id", edinetResources.GetDocumentByType)

	s.app.Post("/documents", docResources.StoreByTerm)

	s.app.Get("/documents", docResources.FindByTerm)
	s.app.Get("/documents/:id", docResources.GetDocument)
	s.app.Get("/companies", companyResources.FindCompanies)
	s.app.Get("/companies/:id", companyResources.FindCompany)

	if s.cfg.Mode == DEVELOPMENT_MODE {
		s.app.Use(pprof.New())
	}
}

func (s *Server) Run() error {
	myLogger.Logger.Info().Msgf("ServerMode: %s", s.cfg.Mode)
	return s.app.Listen(net.JoinHostPort("", conf.LoadServerConfig().Port))
}
