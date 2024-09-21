package micloud

import (
	"errors"
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/config"
	"time"
)

type Config struct {
	MiaccountCookieFilePath string
	MicloudCookieFilePath   string
	Timeout                 time.Duration
	RetryTimes              uint
	NumOfReqInSec           uint
}

func Init(c Config) error {
	if c.MiaccountCookieFilePath == "" {
		return errors.New("no mi account cookie file path")
	} else {
		cookie.MiaccountCookieFilepath = c.MiaccountCookieFilePath
	}
	if c.Timeout != 0 {
		config.Timeout = c.Timeout
	}
	if c.RetryTimes != 0 {
		config.RetryTimes = c.RetryTimes
	}
	if c.NumOfReqInSec != 0 {
		config.NumOfReqInSec = c.NumOfReqInSec
	}
	return nil
}
