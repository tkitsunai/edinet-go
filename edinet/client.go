package edinet

import (
	"context"
	"errors"
	"fmt"
	"github.com/samber/do"
	"github.com/tkitsunai/edinet-go/conf"
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
	Base            string = "https://disclosure.edinet-fsa.go.jp/api/"
	SubscriptionKey string = "Subscription-Key"
)

var (
	apiError = errors.New("[warning] API http status code was not ok")
)

func NewClient(i *do.Injector) (*Client, error) {
	config := do.MustInvoke[*conf.Config](i)
	u, err := url.Parse(Base + v2)
	if err != nil {
		panic("parse base url error")
	}
	return &Client{
		baseUrl:    u,
		httpClient: NewHttpClient(),
		apiKey:     config.ApiKey,
	}, nil
}

func (c *Client) URL() *url.URL {
	baseUrl := *c.baseUrl
	return &baseUrl
}

func (c *Client) RequestDocuments(date core.Date, requestType RequestType) (EdinetDocumentResponse, error) {
	u := c.URL()
	u.Path = path.Join(u.Path, "documents.json")
	q := u.Query()
	q.Set("date", date.String())
	q.Set("type", requestType.String())
	q.Set(SubscriptionKey, c.apiKey)
	u.RawQuery = q.Encode()

	ctx := context.Background()
	res := EdinetDocumentResponse{}
	statusCode, err := c.httpClient.ExecuteGetWithDecodeJson(ctx, u, &res)
	if err != nil {
		return EdinetDocumentResponse{}, err
	}
	if !isSuccessful(statusCode) {
		log.Println("[warning] EDINET-API status code:", statusCode)
		return EdinetDocumentResponse{}, apiError
	}
	return res, nil
}

func (c *Client) RequestDocument(docId core.DocumentId, fileType FileType) (DocumentFile, error) {
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
		return DocumentFile{}, err
	}

	if isSuccessful(response.StatusCode) {
		// response content-type check
		contentType := response.Header.Get("Content-Type")
		var extension string
		switch contentType {
		case "application/octet-stream":
			extension = ".zip"
		case "application/pdf":
			extension = ".pdf"
		case "application/json; charset=utf-8":
			return DocumentFile{}, fmt.Errorf("document not found: ID:%s, type:%s", docId.String(), fileType.String())
		}
		// read binary
		defer response.Body.Close()
		content, err := io.ReadAll(response.Body)
		if err != nil {
			return DocumentFile{}, nil
		}

		res := DocumentFile{
			Name:        fmt.Sprintf("%s_%s", docId.String(), fileType.String()),
			Extension:   extension,
			DocumentId:  docId.String(),
			ContentType: contentType,
			Content:     content,
		}
		return res, nil
	}

	return DocumentFile{}, apiError
}
