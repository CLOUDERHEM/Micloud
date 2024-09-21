package setting

import (
	"errors"
	"fmt"
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/utility/parse"
	"github.com/clouderhem/micloud/utility/request"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	renewalApi = "https://i.mi.com/status/lite/setting?type=AutoRenewal&inactiveTime=10&ts=%v"
)

func Renewal() (string, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(renewalApi, time.Now().UnixMilli()), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Cookie", cookie.GetMicloudCookie())
	_, r, err := request.DoRequest(req)
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(r.StatusCode))
	}

	c := strings.Join(r.Header.Values("Set-Cookie"), ";")
	token := parse.GetValueByKey(c, "serviceToken")
	if token == "" {
		log.Print("serviceToken is empty, do not need to renew")
	}
	return c, nil

}
