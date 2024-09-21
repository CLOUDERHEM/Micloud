package miaccount

import (
	"errors"
	"github.com/clouderhem/micloud/authorizer/cookie"
	usermgr "github.com/clouderhem/micloud/micloud/user"
	"github.com/clouderhem/micloud/utility/parse"
	"github.com/clouderhem/micloud/utility/request"
	"strings"
)

func GetMicloudCookie() (string, error) {
	serviceLoginUrl, err := usermgr.GetLoginUrl()
	if err != nil {
		return "", err
	}
	stsUrl, err := getSTSUrl(serviceLoginUrl)
	if err != nil {
		return "", err
	}
	return getMicloudCookie(stsUrl)
}

func getSTSUrl(serviceLoginUrl string) (string, error) {
	req := request.NewGet(serviceLoginUrl, nil)
	req.Header.Add("Cookie", cookie.GetMiaccountCookie())
	_, resp, err := request.DoRequest(req)
	if err != nil {
		return "", err
	}
	location := resp.Header.Get("Location")
	if location == "" {
		return "", errors.New("no location in service login resp")
	}
	return location, nil
}

func getMicloudCookie(stsUrl string) (string, error) {
	req := request.NewGet(stsUrl, nil)
	req.Header.Set("Cookie", cookie.GetMiaccountCookie())
	_, resp, err := request.DoRequest(req)
	if err != nil {
		return "", err
	}
	values := resp.Header.Values("Set-Cookie")
	if len(values) == 0 {
		return "", errors.New("no cookies in sts resp")
	}
	return parse.TidyKvs(strings.Join(values, ";")), nil
}
