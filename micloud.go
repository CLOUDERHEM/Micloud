package micloud

import (
	"errors"
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/config"
	"time"
)

type Config struct {
	CookieFilepath string
	Timeout        time.Duration
	RetryTimes     uint
	NumOfReqInSec  uint
}

func Init(c Config) error {
	if c.CookieFilepath == "" {
		return errors.New("no cookie file path")
	} else {
		cookie.SetCookieFilepath(c.CookieFilepath)
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
