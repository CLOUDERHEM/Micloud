package miaccount

import (
	"github.com/clouderhem/micloud/authorizer/cookie"
	"testing"
)

func TestGetMicloudCookie(t *testing.T) {
	cookie.MiaccountCookieFilepath = "/misync/.miaccount_cookie"
	micloudCookie, err := GetMicloudCookie()
	if err != nil {
		t.Error(err)
	}
	t.Log(micloudCookie)
}
