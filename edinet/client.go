package edinet

import (
	"context"
	"errors"
	"fmt"
	"github.com/tkitsunai/edinet-go/core"
	"io"
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
	v2              string = "v2"
	v1              string = "v1"
	Base            string = "https://disclosure.edinet-fsa.go.jp/api/"
	SubscriptionKey string = "Subscription-Key"
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
	q.Set(SubscriptionKey, c.apiKey)
	u.RawQuery = q.Encode()

	ctx := context.Background()
	res := DocumentListResponse{}
	statusCode, err := c.httpClient.ExecuteGetWithDecodeJson(ctx, u, &res)
	if err != nil {
		return nil, err
	}
	if !isSuccessful(statusCode) {
		log.Println("[warning] EDINET-API status code:", statusCode)
		return nil, apiError
	}
	return &res, nil
}

func (c *Client) RequestDocument(docId DocumentId, fileType FileType) (File, error) {
	u := c.URL()
	u.Path = path.Join(u.Path, "documents", docId.String())
	q := u.Query()
	q.Set("type", fileType.String())
	q.Set(SubscriptionKey, c.apiKey)
	u.RawQuery = q.Encode()

	ctx := context.Background()
	req, _ := c.httpClient.NewRequest(ctx, "GET", u, "", nil)
	response, err := c.httpClient.client.Do(req)
	if err != nil {
		return File{}, err
	}

	if isSuccessful(response.StatusCode) {
		// response content-type check
		contentType := response.Header.Get("Content-Type")
		fmt.Println(contentType)
		var extension string
		switch contentType {
		case "application/octet-stream":
			extension = ".zip"
		case "application/pdf":
			extension = ".pdf"
		case "application/json; charset=utf-8":
			return File{}, fmt.Errorf("document not found: %s", docId.String())
		}
		// read binary
		defer response.Body.Close()
		content, err := io.ReadAll(response.Body)
		if err != nil {
			return File{}, nil
		}

		res := File{
			Name:       docId.String(),
			Extension:  extension,
			DocumentId: docId,
			Content:    content,
		}
		return res, nil
	}

	return File{}, nil
}
