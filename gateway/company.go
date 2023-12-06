package gateway

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/datastore"
	"github.com/tkitsunai/edinet-go/port"
)

type Company struct {
	db datastore.Driver
}

func NewCompany(i *do.Injector) (port.Company, error) {
	db := do.MustInvoke[datastore.Driver](i)
	return &Company{
		db: db,
	}, nil
}

func (c *Company) FindById(id core.EdinetCode) (core.Company, error) {
	foundData, err := c.db.FindByKey(datastore.CompanyTable, id.String())
	if err != nil {
		return core.Company{}, err
	}
	return decode[core.Company](foundData)
}

func (c *Company) Find() (core.Companies, error) {
	companies, err := c.db.FindAll(datastore.CompanyTable)
	if err != nil {
		return nil, err
	}

	res := make(core.Companies, len(companies))
	for i, company := range companies {
		d, err := decode[core.Company](company)
		if err != nil {
			return nil, err
		}
		res[i] = d
	}

	return res, nil
}

func (c *Company) Store(company core.Company) error {
	// TODO データが見つかった場合
	{
		foundData, err := c.db.FindByKey(datastore.CompanyTable, company.EdinetCode.String())
		if err != nil {
			// nodata
		}
		_, err = decode[core.Company](foundData)
		if err != nil {
			// error decode
			return err
		}
	}
	// データ投入
	return c.db.Update(datastore.CompanyTable, company.EdinetCode.String(), company)
}

func (c *Company) StoreAll(companies core.Companies) error {
	storeData := make(map[string]interface{})
	for _, company := range companies {
		if company.EdinetCode.String() == "" {
			continue
		}
		foundCompany, err := c.FindById(company.EdinetCode)
		// not found
		if err != nil {
			storeData[company.EdinetCode.String()] = company
			continue
		}
		// found
		for k, v := range company.Docs {
			foundCompany.Docs[k] = v
		}
		foundCompany.Name = company.Name
		storeData[company.EdinetCode.String()] = foundCompany
	}
	return c.db.Batch(datastore.CompanyTable, storeData)
}
