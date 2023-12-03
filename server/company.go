package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/usecase"
	"net/http"
)

type Company struct {
	companyUsecase *usecase.Company
}

func NewCompany(i *do.Injector) (*Company, error) {
	companyUsecase := do.MustInvoke[*usecase.Company](i)
	return &Company{companyUsecase: companyUsecase}, nil
}

func (c *Company) FindCompany(ctx *fiber.Ctx) error {
	companyId := ctx.Params("id")

	id := core.EdinetCode(companyId)
	foundCompany, err := c.companyUsecase.FindById(id)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"status":  http.StatusNotFound,
			"message": fmt.Sprintf("company not found ID:%s", companyId),
		})
	}

	return ctx.Status(http.StatusOK).JSON(CompanyJson{
		EdinetCode: foundCompany.ECode.String(),
		Name:       foundCompany.Name.String(),
	})
}

func (c *Company) FindCompanies(ctx *fiber.Ctx) error {
	p := TermParams{}
	err := ctx.QueryParser(&p)
	if err != nil {
		return err
	}

	companies, err := c.companyUsecase.Find()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	results := make([]CompanyJson, len(companies))
	for i, company := range companies {
		s := CompanyJson{
			company.ECode.String(),
			company.Name.String(),
		}
		results[i] = s
	}

	return ctx.Status(http.StatusOK).JSON(struct {
		Companies []CompanyJson `json:"companies"`
	}{
		Companies: results,
	})
}

type CompanyJson struct {
	EdinetCode string `json:"edinetCode"`
	Name       string `json:"name"`
}
