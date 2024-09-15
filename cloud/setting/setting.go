package setting

import (
	"io.github.clouderhem.micloud/cloud/setting/detail"
	"io.github.clouderhem.micloud/cloud/setting/device"
	"io.github.clouderhem.micloud/cloud/setting/renewal"
)

// ListDevices list all devices logged in with current xiaomi account
func ListDevices() ([]device.Device, error) {
	return device.ListDevices()
}

// GetDetail get micloud service detail
func GetDetail() (detail.Detail, error) {
	return detail.GetAllDetail()
}

// Renewal renew cookie
func Renewal() (string, error) {
	cookie, err := renewal.Renewal()
	if err != nil {
		return "", err
	}
	return cookie, err
}
