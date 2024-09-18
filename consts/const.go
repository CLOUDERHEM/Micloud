package consts

import "time"

const (
	DefaultTimeout     = time.Second * 8
	DefaultRetryTimes  = 2
	DefaultReqNumInSec = 10
)

const (
	MicloudCookieFilePath = "/micloud/.micloud_cookie"
)
