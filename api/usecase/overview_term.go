package usecase

import (
	"github.com/tkitsunai/edinet-go/api/domain"
	"github.com/tkitsunai/edinet-go/api/edinet"
	v1 "github.com/tkitsunai/edinet-go/api/edinet/api/v1"
)

type OverviewTerm struct {
	Client *edinet.V1Client
}

func (t OverviewTerm) FindOverviewByTerm(term domain.Term) (v1.DocumentContentResponses, error) {

	//days := term.DayDuration()
	//
	//for _, day := range days {
	//	res, err := t.Client.RequestDocumentList(day)
	//}

	panic("implemented me")
}
