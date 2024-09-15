package device

import "testing"

func TestListDevices(t *testing.T) {
	devices, err := ListDevices()
	if err != nil {
		t.Error(err)
	}
	if len(devices) == 0 {
		t.Error("no devices found")
	}
}
