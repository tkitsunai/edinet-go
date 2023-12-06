package edinet

import (
	"context"
	"encoding/json"
	"github.com/tkitsunai/edinet-go/logger"
	"io"
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

func isSuccessful(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}

func (r *HttpClient) NewRequest(
	ctx context.Context,
	method string,
	u *url.URL,
	contentType string,
	body io.Reader,
) (*http.Request, error) {
	var requestBody string
	if body != nil {
		reqBody, _ := io.ReadAll(body)
		requestBody = string(reqBody)
		logger.Logger.Debug().Msgf("Request body:%s", requestBody)
	}
	logger.Logger.Debug().Msgf("URL:%s", u.String())
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
	req, err := r.NewRequest(ctx, http.MethodGet, u, "application/json", nil)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	logger.Logger.Debug().Msgf("URL:%s", u.String())
	res, err := r.client.Do(req)
	if err != nil {
		return res.StatusCode, err
	}

	if isSuccessful(res.StatusCode) {
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
	req, err := r.NewRequest(ctx, http.MethodGet, url, "application/json", nil)
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
	return io.ReadAll(resp.Body)
}
