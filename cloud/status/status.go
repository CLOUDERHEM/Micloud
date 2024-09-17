package statusmgr

import (
	"io.github.clouderhem.micloud/cloud/status/detail"
	"io.github.clouderhem.micloud/cloud/status/setting"
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
