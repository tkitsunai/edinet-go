package core

type Companies []Company

type CompanyName string

func (c CompanyName) String() string {
	return string(c)
}

type EdinetCode string

func (e EdinetCode) String() string {
	return string(e)
}

type Company struct {
	ECode EdinetCode
	Name  CompanyName
}
