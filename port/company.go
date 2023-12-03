package port

import "github.com/tkitsunai/edinet-go/core"

type Company interface {
	FindById(id core.EdinetCode) (core.Company, error)
	Find() (core.Companies, error)
	Store(company core.Company) error
	StoreAll(companies core.Companies) error
}
