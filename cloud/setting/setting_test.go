package setting

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

func TestGetDetail(t *testing.T) {
	detail, err := GetDetail()
	if err != nil {
		t.Error(err)
	}
	if detail.SettingType == "" {
		t.Error("setting type is empty")
	}
}

func TestRenewal(t *testing.T) {
	renewal, err := Renewal()
	if err != nil {
		t.Error(err)
	}
	if renewal == "" {
		t.Error("no renewal found")
	}
}
