package usermgr

import "testing"

func TestGetLoginUrl(t *testing.T) {
	url, err := GetLoginUrl()
	if err != nil {
		t.Error(err)
	}
	t.Log(url)
}
