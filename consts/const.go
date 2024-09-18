package consts

import (
	"net/url"
	"time"
)

const (
	DefaultTimeout     = time.Second * 8
	DefaultRetryTimes  = 2
	DatabaseDir        = "/micloud"
	DefaultReqNumInSec = 10
)

var (
	MicloudCookieFilePath = ""
)

func init() {
	MicloudCookieFilePath, _ = url.JoinPath(DatabaseDir, ".micloud_cookie")
}
