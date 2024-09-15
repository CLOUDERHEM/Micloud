package device

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
	devicesApi = "https://i.mi.com/passport/user/all/devices?locale=zh_CN?ts="
)

func ListDevices() ([]Device, error) {
	req, err := http.NewRequest(http.MethodGet, devicesApi+fmt.Sprint(time.Now().UnixMilli()), nil)
	if err != nil {
		return nil, err
	}

	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Devices](body)
	if err != nil {
		return nil, err
	}
	if resp.Code != consts.Success {
		return nil, errors.New(resp.Description)
	}
	return resp.Data.List, nil
}
