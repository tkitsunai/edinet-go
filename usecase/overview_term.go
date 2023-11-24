package usecase

import (
	"fmt"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"sync"
	"time"
)

type OverviewTerm struct {
	Client *edinet.Client
}

func NewOverviewTerm(apiKey string) *OverviewTerm {
	return &OverviewTerm{
		Client: edinet.NewClient(apiKey),
	}
}

func (t *OverviewTerm) FindOverviewByDate(date core.FileDate) ([]*edinet.DocumentListResponse, error) {
	var results []*edinet.DocumentListResponse

	doc, err := t.Client.RequestDocumentList(date)

	results = append(results, doc)

	return results, err
}

func (t *OverviewTerm) FindOverviewByTerm(term core.Term) ([]*edinet.DocumentListResponse, []error) {
	days := term.DayDuration()

	var errorsPack []error
	var results []*edinet.DocumentListResponse

	var mutex = &sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, day := range days {
		wg.Add(1)
		mutex.Lock()
		go func(day time.Time) {
			defer wg.Done()
			res, err := t.Client.RequestDocumentList(core.NewFileDate(day))
			if err != nil {
				errorsPack = append(errorsPack, err)
			}
			results = append(results, res)
			mutex.Unlock()
		}(day)
	}
	wg.Wait()

	fmt.Println("Day Size: ", len(days))
	fmt.Println("Response Success Size: ", len(results))
	fmt.Println("Response Error Size: ", len(errorsPack))

	return results, errorsPack
}