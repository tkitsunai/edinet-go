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

func NewDocumentsResource(overview *usecase.Overview, docUsecase *usecase.Document) *Documents {
	return &Documents{
		overviewUsecase: overview,
		documentUsecase: docUsecase,
	}
}

func (d *Documents) GetDocumentsByDate(ctx *fiber.Ctx) error {
	p := FileDateParam{}
	err := ctx.ParamsParser(&p)

	if err != nil {
		return err
	}

	overviews, err := d.overviewUsecase.FindOverviewByDate(core.FileDate(p.Date))

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	res := make([]edinet.DocumentListResponse, 0, len(overviews))
	for idx, overview := range overviews {
		res[idx] = *overview
	}

	return ctx.JSON(edinet.DocumentListResponses{
		List: res,
	})
}

func (d *Documents) GetDocumentsByTerm(ctx *fiber.Ctx) error {
	p := FileTermParams{}
	err := ctx.QueryParser(&p)

	if err != nil {
		return err
	}

	from := core.FileDate(p.From)
	to := core.FileDate(p.To)
	_, err = from.Validate()
	_, err = to.Validate()

	term := core.NewTerm(from, to)

	overviews, errs := d.overviewUsecase.FindOverviewByTerm(term)

	if errs != nil && len(errs) > 0 {
		return ctx.Status(http.StatusInternalServerError).SendString(errs[0].Error())
	}

	res := make([]edinet.DocumentListResponse, len(overviews))
	for idx, overview := range overviews {
		res[idx] = *overview
	}

	return ctx.JSON(edinet.DocumentListResponses{
		List: res,
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

	document, err := d.documentUsecase.FindContentById(documentId, fileType)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// download file
	ctx.Set("Content-Disposition", "attachment; filename="+document.NameWithExtension())
	return ctx.Send(document.Content)
}

type FileDateParam struct {
	Date string
}

type FileTermParams struct {
	From string `query:"from"`
	To   string `query:"to"`
}
