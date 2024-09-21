package user

import (
	"errors"
	"fmt"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"net/http"
	"time"
)

const (
	loginUrl = "https://i.mi.com/api/user/login?&followUp=https%3A%2F%2Fi.mi.com%2F&_locale=zh_CN&ts="
)

func GetLoginUrl() (string, error) {
	req := request.NewGet(loginUrl+fmt.Sprintf("%v", time.Now().UnixMilli()), nil)
	body, r, err := request.DoRequest(req)
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(r.StatusCode))
	}

	data, err := validate.Validate[LoginUrl](r, body)
	if err != nil {
		return "", err
	}
	return data.LoginUrl, err
}
