package statusmgr

import (
	"github.com/clouderhem/micloud/micloud/status/detail"
	"github.com/clouderhem/micloud/micloud/status/setting"
)

// GetDetail get micloud service detail
func GetDetail() (detail.Detail, error) {
	return detail.GetAllDetail()
}

// Renewal renew cookie
func Renewal() (string, error) {
	cookie, err := setting.Renewal()
	if err != nil {
		return "", err
	}
	return cookie, err
}
