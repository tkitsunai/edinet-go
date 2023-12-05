package adapter

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
)

type CompanyConverter struct {
}

func (c *CompanyConverter) UniqueCompanies(results []edinet.Result) ([]core.Company, error) {
	// create unique company list
	unique := make(map[string]core.Company)
	for _, result := range results {
		company := core.Company{
			EdinetCode: core.EdinetCode(result.EdinetCode),
			Name:       core.CompanyName(result.FilerName),
			Docs:       make(map[string]core.Document),
		}
		unique[company.EdinetCode.String()] = company
	}
	// doc code
	for _, result := range results {
		documentId := result.DocID
		if len(documentId) == 0 {
			continue
		}
		edinetCode := result.EdinetCode
		if len(edinetCode) == 0 {
			continue
		}
		// find company
		if company, ok := unique[edinetCode]; ok {
			// append doc id
			document := core.Document{Id: core.DocumentId(documentId)}
			company.Docs[documentId] = document
		}
	}

	r := make([]core.Company, 0, len(unique))
	for _, v := range unique {
		r = append(r, v)
	}

	return r, nil
}

func NewCompanyConverter(_ *do.Injector) (port.CompanyConverter, error) {
	return &CompanyConverter{}, nil
}
