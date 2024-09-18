package validate

import (
	"errors"
	"fmt"
	"github.com/clouderhem/micloud/consts"
	"github.com/clouderhem/micloud/utility/response"
	"net/http"
)

// Validate validate response body, and return data
func Validate[T any](r *http.Response, body []byte) (data T, err error) {
	if r.StatusCode != http.StatusOK {
		var t T
		return t, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[T](body)
	if err != nil {
		return resp.Data, err
	}

	if resp.Code != consts.Success {
		return resp.Data, errors.New(fmt.Sprintf("%v:%v", resp.Description, resp.Result))
	}
	return resp.Data, nil
}
