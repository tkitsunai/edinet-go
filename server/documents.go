package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/usecase"
	"net/http"
)

type Documents struct {
	overviewUsecase *usecase.Overview
	documentUsecase *usecase.Document
}

func NewDocumentsResource(
	overview *usecase.Overview,
	docUsecase *usecase.Document,
) *Documents {
	return &Documents{
		overviewUsecase: overview,
		documentUsecase: docUsecase,
	}
}

func (d *Documents) StoreDocumentsByTerm(ctx *fiber.Ctx) error {
	p := TermParams{}
	err := ctx.QueryParser(&p)
	if err != nil {
		return err
	}

	term := core.NewTerm(core.Date(p.From), core.Date(p.To))

	err = d.overviewUsecase.StoreByTerm(term)
	if err != nil {
		return err
	}
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"status": "ok",
	})
}

func (d *Documents) GetDocumentsByTerm(ctx *fiber.Ctx) error {
	p := TermParams{}
	err := ctx.QueryParser(&p)

	if err != nil {
		return err
	}

	term := core.NewTerm(core.Date(p.From), core.Date(p.To))

	overviews, err := d.overviewUsecase.FindOverviewByTerm(term, p.Refresh)

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

func (d *Documents) GetDocument(ctx *fiber.Ctx) error {
	did := ctx.Params("id")
	documentId := edinet.DocumentId(did)

	query := ctx.Query("type")
	fileType := edinet.NewFileTypeByName(query)

	if fileType == edinet.Unknown {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "request parameter invalid. parameter name 'type' is required",
		})
	}

	document, err := d.documentUsecase.FindContent(documentId, fileType)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	ctx.Set("Content-Disposition", "attachment; filename="+document.NameWithExtension())
	return ctx.Send(document.Content)
}

type FileDateParam struct {
	Date string
}

type TermParams struct {
	From    string `query:"from"`
	To      string `query:"to"`
	Refresh bool   `query:"refresh"`
}
