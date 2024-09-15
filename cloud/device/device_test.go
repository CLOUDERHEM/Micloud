package device

import (
	"fmt"
	"testing"
)

func TestListDevices(t *testing.T) {
	devices, err := ListDevices()
	if err != nil {
		t.Error(err)
	}
	if len(devices) == 0 {
		t.Error("no devices found")
	}
}

func TestGetDeviceStatus(t *testing.T) {
	status, err := GetDeviceStatus()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(status)
}
