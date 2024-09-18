package devicemgr

import (
	"github.com/clouderhem/micloud/micloud/device/device"
	"github.com/clouderhem/micloud/micloud/device/status"
)

// ListDevices list all devices logged in with current xiaomi account
func ListDevices() ([]device.Device, error) {
	return device.ListDevices()
}

// GetDeviceStatus list all devices' status including position info
func GetDeviceStatus() (status.Status, error) {
	return status.GetDeviceStatus()
}
