package usecase

import (
	"github.com/hashicorp/go-multierror"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/logger"
	"github.com/tkitsunai/edinet-go/port"
	"sync"
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

func (t *Overview) FindOverviewByTerm(term core.Term, refresh bool) ([]edinet.EdinetDocumentResponse, error) {
	dateRange := term.GetDateRange()
	//var results []edinet.EdinetDocumentResponse
	var mu = sync.Mutex{}
	var meg multierror.Group

	results := make([]edinet.EdinetDocumentResponse, len(dateRange))
	for i, date := range dateRange {
		runAsync := func(idx int, date core.Date) {
			mu.Lock()
			defer mu.Unlock()
			meg.Go(func() error {
				// data store exists, using stored data.
				// If refresh mode is on, the data is not retrieved from the data store.
				if !refresh {
					if store, err := t.ovPort.GetByStore(date); err == nil {
						logger.Logger.Info().Msgf("find stored data. %s", date)
						results[idx] = store
						return nil
					}
				}

				res, err := t.ovPort.Get(date)
				if err != nil {
					return err
				}
				results[idx] = res
				return nil
			})
		}
		runAsync(i, date)
	}
	waitedError := meg.Wait()
	if waitedError != nil {
		return nil, waitedError.ErrorOrNil()
	}

	logger.Logger.Info().Msgf("Day Size: %d", len(dateRange))
	logger.Logger.Info().Msgf("Response Success Size: %d", len(results))

	for _, data := range results {
		logger.Logger.Info().Msgf("Day: %s Document Results Size: %d", data.Metadata.Parameter.Date, len(data.Results))
	}

	return results, nil
}
