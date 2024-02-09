package core

type Companies []Company
type CompanyName string
type EdinetCode string

type Company struct {
	EdinetCode EdinetCode
	Name       CompanyName

	Docs map[string]Document
}

type DocumentId string

type Document struct {
	Id DocumentId
}

func (c CompanyName) String() string {
	return string(c)
}

func (e EdinetCode) String() string {
	return string(e)
}

func (d DocumentId) String() string {
	return string(d)
}

func (d DocumentId) Valid() bool {
	return len(string(d)) > 0
}

func (e EdinetCode) Valid() bool {
	return len(string(e)) > 0
}

func (c Companies) FilterUnknownEdinetCode() Companies {
	var res Companies
	for _, company := range c {
		if company.EdinetCode.Valid() {
			res = append(res, company)
		}
	}
	return res
}
