package detail

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
	"time"
)

const (
	detailApi = "https://i.mi.com/status/lite/alldetail?ts="
)

func GetAllDetail() (Detail, error) {
	req, err := http.NewRequest(http.MethodGet, detailApi+fmt.Sprint(time.Now().UnixMilli()), nil)
	if err != nil {
		return Detail{}, nil
	}

	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return Detail{}, nil
	}

	if r.StatusCode != http.StatusOK {
		return Detail{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Detail](body)
	if err != nil {
		return Detail{}, nil
	}
	if resp.Code != consts.Success {
		return Detail{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
