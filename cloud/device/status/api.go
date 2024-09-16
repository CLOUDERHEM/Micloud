package status

import (
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/validate"
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

	return validate.Validate[Status](r, body)
}
