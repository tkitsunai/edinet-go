package core

import (
	"io"
	"net/http"
)

var instance *HttpClient

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	if instance == nil {
		instance = &HttpClient{
			client: http.DefaultClient,
		}
	}
	return instance
}

func (r *HttpClient) ExecuteGetJsonType(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", JSONContentType.String())
	res, err := (http.DefaultClient).Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *HttpClient) Request(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := (http.DefaultClient).Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
