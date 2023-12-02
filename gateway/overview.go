package gateway

import (
	"bytes"
	"encoding/gob"
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
	c := do.MustInvoke[*edinet.Client](i)
	db := do.MustInvoke[datastore.Driver](i)
	return &Overview{
		c:  c,
		db: db,
	}, nil
}

func (o *Overview) GetByStore(date core.Date) (edinet.EdinetDocumentResponse, error) {
	// メタデータの取得
	findMetaData, err := o.db.FindByKey("metadata", date.String())
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

func decode[T edinet.Result | edinet.Metadata](data []byte) (T, error) {
	var emp T
	var result T
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&result)
	if err != nil {
		return emp, err
	}
	return result, nil
}

func (o *Overview) Get(date core.Date) (edinet.EdinetDocumentResponse, error) {
	// data persistent from edinet-data
	results, err := o.c.RequestDocumentList(date)
	if err != nil {
		return edinet.EdinetDocumentResponse{}, err
	}

	// all data saving
	batchingData := make(map[string]interface{})
	for _, overview := range results.Results {
		batchingData[overview.DocID] = overview
	}

	storeErr := o.db.Batch(results.Metadata.Parameter.Date, batchingData)
	err = o.db.Update("metadata", results.Metadata.Parameter.Date, results.Metadata)
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
