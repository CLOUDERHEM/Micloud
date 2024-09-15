package setting

import (
	"io.github.clouderhem.micloud/cloud/setting/device"
)

// ListDevices list all devices logged in with current xiaomi account
func ListDevices() ([]device.Device, error) {
	return device.ListDevices()
}
