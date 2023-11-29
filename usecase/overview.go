package usecase

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
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

func (t *Overview) FindOverviewByTerm(term core.Term) ([]*edinet.DocumentListResponse, error) {
	dateRange := term.GetDateRange()
	var results []*edinet.DocumentListResponse

	var resultsLock = sync.Mutex{}
	var meg multierror.Group
	var wg sync.WaitGroup
	for _, date := range dateRange {
		runAsync := func(date time.Time) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				meg.Go(func() error {
					res, err := t.Client.RequestDocumentList(core.NewFileDate(date))
					if err != nil {
						return err
					}
					resultsLock.Lock()
					results = append(results, res)
					resultsLock.Unlock()
					return nil
				})
			}()
		}
		runAsync(date)
	}
	wg.Wait()
	waitedError := meg.Wait()

	var err error

	if waitedError != nil {
		for _, goerr := range waitedError.Errors {
			err = multierror.Append(err, goerr)
		}
	}

	logger.Logger.Info().Msg(fmt.Sprintf("Day Size: %d", len(dateRange)))
	logger.Logger.Info().Msg(fmt.Sprintf("Response Success Size: %d", len(results)))

	return results, err
}
