package recorder

import (
	"testing"
)

func TestListRecorders(t *testing.T) {
	recorders, err := ListRecorders(0, 10)
	if err != nil {
		t.Error(err)
	}
	if len(recorders) == 0 {
		t.Error("no recorder")
	}
}

func TestGetRecorder(t *testing.T) {
	url, err := GetRecorderUrl("43552880029221120")
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("url is empty")
	}
}

func TestDeleteRecorder(t *testing.T) {
	err := DeleteRecorder("43410534532923776")
	if err != nil {
		t.Error(err)
	}
}
