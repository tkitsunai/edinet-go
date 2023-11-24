package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/usecase"
	"net/http"
)

type Documents struct {
	useCase *usecase.OverviewTerm
}

func NewDocumentsResource(apiKey string) *Documents {
	return &Documents{
		useCase: usecase.NewOverviewTerm(apiKey),
	}
}

func (d *Documents) GetDocuments(ctx *fiber.Ctx) error {
	p := FileDateParam{}
	err := ctx.ParamsParser(&p)

	if err != nil {
		return err
	}

	overviews, err := d.useCase.FindOverviewByDate(core.FileDate(p.Date))

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
	p := FileTermParam{}
	err := ctx.QueryParser(&p)

	if err != nil {
		return err
	}

	from := core.FileDate(p.From)
	to := core.FileDate(p.To)
	_, err = from.Validate()
	_, err = to.Validate()

	term := core.NewTerm(from, to)

	overviews, errs := d.useCase.FindOverviewByTerm(term)

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

type FileDateParam struct {
	Date string
}

type FileTermParam struct {
	From string `query:"from"`
	To   string `query:"to"`
}
