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
