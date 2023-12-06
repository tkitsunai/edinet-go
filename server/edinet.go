package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/usecase"
	"net/http"
)

type EdinetRaw struct {
	document *usecase.Document
	overview *usecase.Overview
}

func (r *EdinetRaw) GetMetaDataByDate(ctx *fiber.Ctx) error {
	p := EdinetParam{}
	err := ctx.QueryParser(&p)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	requestType := edinet.RequestType(p.Type)
	result, err := r.overview.FindByDateAndType(core.Date(p.Date), requestType)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(result)
}

func (r *EdinetRaw) GetDocumentByType(ctx *fiber.Ctx) error {
	docId := ctx.Params("id")
	if len(docId) == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": fmt.Errorf("document id path not match").Error(),
		})
	}
	id := core.DocumentId(docId)

	query := ctx.Query("type")
	fileType := edinet.NewFileTypeByName(query)
	document, err := r.document.FindContent(id, fileType)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	ctx.Set("Content-Disposition", "attachment; filename="+document.NameWithExtension())
	return ctx.Status(http.StatusOK).Send(document.Content)
}

func NewEdinetRaw(i *do.Injector) (*EdinetRaw, error) {
	return &EdinetRaw{
		document: do.MustInvoke[*usecase.Document](i),
		overview: do.MustInvoke[*usecase.Overview](i),
	}, nil
}

type EdinetParam struct {
	Date string `query:"date"`
	Type int    `query:"type"`
	// TODO 今は不要
	// Subscription-Key
}
