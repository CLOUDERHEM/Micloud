package device

import (
	"io.github.clouderhem.micloud/cloud/device/device"
)

// ListDevices list all devices logged in with current xiaomi account
func ListDevices() ([]device.Device, error) {
	return device.ListDevices()
}
