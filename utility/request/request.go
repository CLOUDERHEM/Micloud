package request

import (
	"github.com/clouderhem/micloud/consts"
	"io"
	"net/http"
)

type UrlQuery struct {
	Key   string
	Value string
}

func DoRequest(req *http.Request) (body []byte, resp *http.Response, err error) {
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	client.Timeout = consts.DefaultTimeout
	defer client.CloseIdleConnections()

	resp, err = client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return data, resp, nil
}

func NewGet(url string, queries []UrlQuery) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	for _, kv := range queries {
		q.Add(kv.Key, kv.Value)
	}
	req.URL.RawQuery = q.Encode()

	return req
}
