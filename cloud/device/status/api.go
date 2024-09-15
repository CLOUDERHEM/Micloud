package status

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
	"time"
)

const (
	deviceStatusApi = "https://i.mi.com/find/device/full/status?ts=%v"
)

func GetDeviceStatus() (Status, error) {
	req := request.NewGet(fmt.Sprintf(deviceStatusApi, time.Now().UnixMilli()), nil)
	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return Status{}, err
	}

	if r.StatusCode != http.StatusOK {
		return Status{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Status](body)
	if err != nil {
		return Status{}, err
	}
	if resp.Code != consts.Success {
		return Status{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
