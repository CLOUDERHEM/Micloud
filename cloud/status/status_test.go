package statusmgr

import "testing"

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
		t.Error("no cookie found")
	}
}
