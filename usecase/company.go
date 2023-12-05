package usecase

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/port"
)

type Company struct {
	companyPort port.Company
}

func NewCompany(i *do.Injector) (*Company, error) {
	companyPort := do.MustInvoke[port.Company](i)
	return &Company{companyPort: companyPort}, nil
}

func (c *Company) FindById(id core.EdinetCode) (core.Company, error) {
	return c.companyPort.FindById(id)
}

func (c *Company) Find() (core.Companies, error) {
	findCompanies, err := c.companyPort.Find()
	if err != nil {
		return nil, err
	}
	return findCompanies, nil
}
