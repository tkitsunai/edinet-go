package usecase

import (
	"fmt"
	"github.com/tkitsunai/edinet-go/api/domain"
	"github.com/tkitsunai/edinet-go/api/edinet"
	v1 "github.com/tkitsunai/edinet-go/api/edinet/api/v1"
)

type OverviewTerm struct {
	Client *edinet.V1Client
}

func NewOverviewTerm(client *edinet.V1Client) *OverviewTerm {
	if client == nil {
		client = edinet.NewV1Client()
	}

	return &OverviewTerm{
		Client: client,
	}
}

func (t OverviewTerm) FindOverviewByTerm(term domain.Term) ([]*v1.DocumentListResponse, []error) {
	days := term.DayDuration()

	var errorsPack []error
	var results []*v1.DocumentListResponse

	for _, day := range days {
		res, err := t.Client.RequestDocumentList(v1.NewFileDate(day))

		if err != nil {
			errorsPack = append(errorsPack, err)
			continue
		}

		results = append(results, res)
	}

	fmt.Println("Day Size: ", len(days))
	fmt.Println("Response Success Size: ", len(results))
	fmt.Println("Response Error Size: ", len(errorsPack))

	return results, errorsPack
}
