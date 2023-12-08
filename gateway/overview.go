package gateway

import (
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/datastore"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/logger"
	"github.com/tkitsunai/edinet-go/port"
)

type Overview struct {
	c  *edinet.Client
	db datastore.Driver
}

func NewOverview(i *do.Injector) (port.Overview, error) {
	return &Overview{
		c:  do.MustInvoke[*edinet.Client](i),
		db: do.MustInvoke[datastore.Driver](i),
	}, nil
}

func (o *Overview) GetRaw(date core.Date, requestType edinet.RequestType) (edinet.EdinetDocumentResponse, error) {
	return o.c.RequestDocuments(date, requestType)
}

func (o *Overview) GetByStore(date core.Date) (edinet.EdinetDocumentResponse, error) {
	// メタデータの取得
	findMetaData, err := o.db.FindByKey(datastore.MetaDataTable, date.String())
	if err != nil {
		return edinet.EdinetDocumentResponse{}, err
	}
	decodedMetaData, err := decode[edinet.Metadata](findMetaData)
	if err != nil {
		return edinet.EdinetDocumentResponse{}, err
	}

	// 日付データの取得
	foundResults, err := o.db.FindAll(date.String())

	results := make([]edinet.Result, len(foundResults))
	for i, result := range foundResults {
		// decode results
		decodedResult, err := decode[edinet.Result](result)
		if err != nil {
			return edinet.EdinetDocumentResponse{}, err
		}
		results[i] = decodedResult
	}

	return edinet.EdinetDocumentResponse{
		Metadata: decodedMetaData,
		Results:  results,
	}, nil
}

func (o *Overview) Get(date core.Date) (edinet.EdinetDocumentResponse, error) {
	// data persistent from edinet-data
	results, err := o.c.RequestDocuments(date, edinet.MetaDataAndDocuments)
	if err != nil {
		return edinet.EdinetDocumentResponse{}, err
	}

	// all data saving
	batchingData := make(map[string]interface{})
	for _, overview := range results.Results {
		batchingData[overview.DocID] = overview
	}

	storeErr := o.db.Batch(results.Metadata.Parameter.Date, batchingData)
	err = o.db.Update(datastore.MetaDataTable, results.Metadata.Parameter.Date, results.Metadata)
	if err != nil {
		// if failed save error
		// TODO requires consideration
		logger.Logger.Error().Msg(err.Error())
	}

	if storeErr != nil {
		// TODO requires consideration
		logger.Logger.Error().Msg(err.Error())
	}

	return results, nil
}
