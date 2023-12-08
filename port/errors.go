package port

import "errors"

var (
	DocumentNotFound = errors.New("document not found")
	CompanyNotFound  = errors.New("company not found")
)
