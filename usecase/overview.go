package usecase

import (
	"fmt"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/logger"
	"sync"
	"time"
)

type Overview struct {
	Client *edinet.Client
}

func NewOverview(client *edinet.Client) *Overview {
	return &Overview{
		Client: client,
	}
}

func (t *Overview) FindOverviewByDate(date core.FileDate) ([]*edinet.DocumentListResponse, error) {
	var results []*edinet.DocumentListResponse

	doc, err := t.Client.RequestDocumentList(date)

	results = append(results, doc)

	return results, err
}

func (t *Overview) FindOverviewByTerm(term core.Term) ([]*edinet.DocumentListResponse, []error) {
	dateRange := term.GetDateRange()

	var errorsPack []error
	var results []*edinet.DocumentListResponse

	var mutex = &sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, date := range dateRange {
		wg.Add(1)
		mutex.Lock()
		go func(date time.Time) {
			defer wg.Done()
			res, err := t.Client.RequestDocumentList(core.NewFileDate(date))
			if err != nil {
				errorsPack = append(errorsPack, err)
			}
			results = append(results, res)
			mutex.Unlock()
		}(date)
	}
	wg.Wait()

	logger.Logger.Info().Msg(fmt.Sprintf("Day Size: %d", len(dateRange)))
	logger.Logger.Info().Msg(fmt.Sprintf("Response Success Size: %d", len(results)))

	return results, errorsPack
}
