package edinet

import (
	"context"
	"errors"
	"github.com/tkitsunai/edinet-go/core"
	"log"
	"net/url"
	"path"
)

type Client struct {
	baseUrl    *url.URL
	httpClient *HttpClient
	apiKey     string
}

const (
	v2   string = "v2"
	v1   string = "v1"
	Base string = "https://disclosure.edinet-fsa.go.jp/api/"
)

var (
	apiError = errors.New("[warning] API http status code was not ok")
)

func NewClient(apiKey string) *Client {
	u, err := url.Parse(Base + v2)
	if err != nil {
		panic("parse base url error")
	}
	return &Client{
		baseUrl:    u,
		httpClient: NewHttpClient(),
		apiKey:     apiKey,
	}
}

func (c *Client) URL() *url.URL {
	copy := *c.baseUrl
	return &copy
}

func (c *Client) RequestDocumentList(date core.FileDate) (*DocumentListResponse, error) {
	u := c.URL()
	u.Path = path.Join(u.Path, "documents.json")
	q := u.Query()
	q.Set("date", date.Format())
	q.Set("type", MetaDataAndDocuments.String())
	q.Set("Subscription-Key", c.apiKey)
	u.RawQuery = q.Encode()

	ctx := context.Background()
	res := DocumentListResponse{}
	statusCode, err := c.httpClient.ExecuteGetWithDecodeJson(ctx, u, &res)
	if err != nil {
		return nil, err
	}
	if !IsOK(statusCode) {
		log.Println("[warning] EDINET-API status code:", statusCode)
		return nil, apiError
	}
	return &res, nil
}

func (c *Client) RequestDocumentContent(requestType RequestType) (*DocumentListResponse, error) {
	panic("TODO")
}
