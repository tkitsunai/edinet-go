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
	ovPort      port.Overview
	companyPort port.Company
	ccPort      port.CompanyConverter
}

func NewOverview(i *do.Injector) (*Overview, error) {
	return &Overview{
		ovPort:      do.MustInvoke[port.Overview](i),
		companyPort: do.MustInvoke[port.Company](i),
		ccPort:      do.MustInvoke[port.CompanyConverter](i),
	}, nil
}

func (o *Overview) StoreByTerm(term core.Term) error {
	dateRange := term.GetDateRange()
	var mutex sync.Mutex

	responses := make([]edinet.EdinetDocumentResponse, len(dateRange))
	wg := sync.WaitGroup{}
	for _, date := range dateRange {
		wg.Add(1)
		go func(date core.Date) error {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			gr, err := o.ovPort.Get(date)
			responses = append(responses, gr)
			if err != nil {
				return err
			}
			return nil
		}(date)
	}
	wg.Wait()

	// 複数の日付にまたがって存在する企業と書類情報
	totalSize := 0
	for _, resultsSet := range responses {
		totalSize += len(resultsSet.Results)
	}
	joinResults := make([]edinet.Result, 0, totalSize)
	for _, resultsSet := range responses {
		joinResults = append(joinResults, resultsSet.Results...)
	}
	companies, err := o.ccPort.UniqueCompanies(joinResults)
	if err != nil {
		return err
	}
	err = o.companyPort.StoreAll(companies)
	if err != nil {
		return err
	}

	return nil
}

func (o *Overview) FindByTerm(term core.Term) ([]edinet.Result, error) {
	dateRange := term.GetDateRange()

	var mu = sync.Mutex{}
	var meg = multierror.Group{}
	var wg sync.WaitGroup

	var results []edinet.Result
	for i, date := range dateRange {
		wg.Add(1)
		go func(idx int, date core.Date) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			meg.Go(func() error {
				storedDocumentResults, err := o.ovPort.GetByStore(date)
				if err != nil {
					return err
				}
				results = append(results, storedDocumentResults.Results...)
				return nil
			})
		}(i, date)
	}
	wg.Wait()
	megErr := meg.Wait()

	if err := megErr.ErrorOrNil(); err != nil {
		return nil, err
	}

	return results, nil
}

func (o *Overview) FindRawByTerm(term core.Term, refresh bool) ([]edinet.EdinetDocumentResponse, error) {
	dateRange := term.GetDateRange()
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
					if store, err := o.ovPort.GetByStore(date); err == nil {
						logger.Logger.Info().Msgf("find stored data. %s", date)
						results[idx] = store
						return nil
					}
				}

				res, err := o.ovPort.Get(date)
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
