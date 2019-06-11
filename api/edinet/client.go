package edinet

import (
	"context"
	"errors"
	v1 "github.com/tkitsunai/edinet-go/api/edinet/api/v1"
	"github.com/tkitsunai/edinet-go/api/edinet/driver"
	"log"
	"net/url"
	"path"
)

type V1Client struct {
	baseUrl    *url.URL
	httpClient *driver.Client
}

const (
	v1Base string = "https://disclosure.edinet-fsa.go.jp/api/v1"
)

var (
	apiError = errors.New("[warning] API http status code was not ok")
)

func NewV1Client() *V1Client {
	u, err := url.Parse(v1Base)
	if err != nil {
		panic("parse base url error")
	}
	return &V1Client{
		baseUrl:    u,
		httpClient: driver.NewClient(),
	}
}

func (c *V1Client) URL() *url.URL {
	copy := *c.baseUrl
	return &copy
}

func (c V1Client) RequestDocumentList(date v1.FileDate) (*v1.DocumentListResponse, error) {
	u := c.URL()
	u.Path = path.Join(u.Path, "documents.json")
	q := u.Query()
	q.Set("date", date.Format())
	q.Set("type", v1.MetaDataAndDocuments.String())
	u.RawQuery = q.Encode()

	ctx := context.Background() // FIXME tobe argument
	res := v1.DocumentListResponse{}
	statusCode, err := c.httpClient.ExecuteGetWithDecodeJson(ctx, u, &res)
	if err != nil {
		return nil, err
	}
	if !driver.IsOK(statusCode) {
		log.Println("[warning] EDINET-API status code:", statusCode)
		return nil, apiError
	}
	return &res, nil
}

func (c V1Client) RequestDocumentContent(requestType v1.RequestType) error {
	panic("implement me")
}
