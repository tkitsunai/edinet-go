package usecase

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/logger"
	"github.com/tkitsunai/edinet-go/port"
	"sync"
	"time"
)

type Overview struct {
	ovPort port.Overview
}

func NewOverview(i *do.Injector) (*Overview, error) {
	ovPort := do.MustInvoke[port.Overview](i)
	return &Overview{
		ovPort: ovPort,
	}, nil
}

func (t *Overview) FindOverviewByDate(date core.FileDate) ([]*edinet.DocumentListResponse, error) {
	var results []*edinet.DocumentListResponse

	doc, err := t.ovPort.Get(date)

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
					fileDate := core.NewFileDate(date)
					// data store exists, using stored data.
					if store, err := t.ovPort.GetByStore(fileDate); err == nil {
						logger.Logger.Info().Msg(fmt.Sprintf("find stored data. %s", fileDate))
						resultsLock.Lock()
						results = append(results, store)
						resultsLock.Unlock()
						return nil
					}

					res, err := t.ovPort.Get(fileDate)
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

	for _, data := range results {
		logger.Logger.Info().Msg(fmt.Sprintf("Day: %s Document Results Size: %d", data.Metadata.Parameter.Date, len(data.Results)))
	}

	return results, err
}
