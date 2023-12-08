package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/usecase"
	"net/http"
)

type Overview struct {
	overviewUC *usecase.Overview
	documentUC *usecase.Document
}

var validate = validator.New()

func NewOverview(
	i *do.Injector,
) (*Overview, error) {
	return &Overview{
		overviewUC: do.MustInvoke[*usecase.Overview](i),
		documentUC: do.MustInvoke[*usecase.Document](i),
	}, nil
}

func (d *Overview) StoreByTerm(ctx *fiber.Ctx) error {
	p := TermParams{}
	err := ctx.BodyParser(&p)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	err = validate.Struct(p)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	term := core.NewTerm(core.Date(p.From), core.Date(p.To))

	err = d.overviewUC.StoreByTerm(term)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func (d *Overview) FindByTerm(ctx *fiber.Ctx) error {
	p := TermParams{}
	err := ctx.QueryParser(&p)

	if err != nil {
		return err
	}

	term := core.NewTerm(core.Date(p.From), core.Date(p.To))

	overviews, err := d.overviewUC.FindByTerm(term, p.Refresh)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	res := make([]edinet.EdinetDocumentResponse, len(overviews))
	for idx, overview := range overviews {
		res[idx] = overview
	}

	return ctx.JSON(edinet.EdinetResponses{
		Items: res,
	})
}

func (d *Overview) GetDocument(ctx *fiber.Ctx) error {
	did := ctx.Params("id")
	documentId := core.DocumentId(did)
	query := ctx.Query("type")
	fileType := edinet.NewFileTypeByName(query)

	if ok := documentId.Valid(); !ok || fileType == edinet.Unknown {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "request parameter invalid. parameter name 'type' is required",
		})
	}

	document, err := d.documentUC.Download(documentId, fileType)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	ctx.Set("Content-Disposition", "attachment; filename="+document.NameWithExtension())
	return ctx.Send(document.Content)
}

type TermParams struct {
	//
	From    string `query:"from" validate:"required"`
	To      string `query:"to" validate:"required"`
	Refresh bool   `query:"refresh"`
}
