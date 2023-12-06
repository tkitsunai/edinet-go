package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/usecase"
)

type AdminResources struct {
	document *usecase.Document
	overview *usecase.Overview
	company  *usecase.Company
}

func NewAdminResources(i *do.Injector) (*AdminResources, error) {
	return &AdminResources{
		document: do.MustInvoke[*usecase.Document](i),
		overview: do.MustInvoke[*usecase.Overview](i),
		company:  do.MustInvoke[*usecase.Company](i),
	}, nil
}

func (a *AdminResources) StoredData(ctx *fiber.Ctx) error {
	return nil
}
