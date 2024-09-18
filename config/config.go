package config

import "time"

var (
	Timeout            = time.Second * 8
	RetryTimes    uint = 2
	NumOfReqInSec uint = 10
)
