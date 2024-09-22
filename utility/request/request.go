package request

import (
	"github.com/clouderhem/micloud/config"
	"io"
	"net/http"
)

type UrlQuery struct {
	Key   string
	Value string
}

func DoRequest(req *http.Request) (body []byte, resp *http.Response, err error) {
	return doRequest(req, true)
}

func DoRequestNotReadBody(req *http.Request) (resp *http.Response, err error) {
	_, resp, err = doRequest(req, false)
	return
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

func doRequest(req *http.Request, readBody bool) (body []byte, resp *http.Response, err error) {
	client := http.DefaultClient
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	client.Timeout = config.Timeout
	defer client.CloseIdleConnections()

	resp, err = client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	if readBody {
		defer resp.Body.Close()
		data, err := io.ReadAll(resp.Body)
		return data, resp, err
	}
	return nil, resp, nil
}
