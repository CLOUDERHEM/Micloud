package renewal

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"net/http"
	"strings"
	"time"
)

const (
	renewalApi = "https://i.mi.com/status/lite/setting?type=AutoRenewal&inactiveTime=10&ts=%v"
)

func Renewal() (cookie string, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(renewalApi, time.Now().UnixMilli()), nil)
	if err != nil {
		return "", err
	}
	_, r, err := authorizer.DoRequest(req)
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(r.StatusCode))
	}

	return strings.Join(r.Header.Values("Set-Cookie"), ";"), nil

}
