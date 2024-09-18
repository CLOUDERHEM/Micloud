package device

import (
	"errors"
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/utility/validate"
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

	data, err := validate.Validate[Devices](r, body)
	if err != nil {
		return nil, err
	}
	return data.List, err
}
