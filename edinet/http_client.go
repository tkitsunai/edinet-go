package edinet

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: http.DefaultClient,
	}
}

func IsOK(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}

func (r *HttpClient) newRequest(
	ctx context.Context,
	method string,
	u *url.URL,
	contentType string,
	body io.Reader,
) (*http.Request, error) {
	log.Println("Request to:", u.String())

	var requestBody string
	if body != nil {
		reqBody, _ := io.ReadAll(body)
		requestBody = string(reqBody)
		log.Println("Request body:", requestBody)
	}

	req, err := http.NewRequest(method, u.String(), strings.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if len(contentType) != 0 {
		req.Header.Set("Content-Type", contentType)
	}

	return req, nil
}

func (r *HttpClient) ExecuteGetWithDecodeJson(
	ctx context.Context,
	u *url.URL,
	out interface{},
) (statusCode int, err error) {
	req, err := r.newRequest(ctx, http.MethodGet, u, "application/json", nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	log.Println("URL:", u.String())
	res, err := r.client.Do(req)
	if err != nil {
		return res.StatusCode, err
	}

	if IsOK(res.StatusCode) {
		err = r.DecodeJsonBody(res, out)
		if err != nil {
			return res.StatusCode, err
		}
	}

	return res.StatusCode, nil
}

func (r *HttpClient) ExecuteGetJsonWithDecodeString(
	ctx context.Context,
	url *url.URL,
) (responseBody string, statusCode int, err error) {
	req, err := r.newRequest(ctx, http.MethodGet, url, "application/json", nil)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	decode, err := r.DecodeByte(res)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	return string(decode), res.StatusCode, nil
}

func (r *HttpClient) DecodeJsonBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (r *HttpClient) DecodeByte(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
