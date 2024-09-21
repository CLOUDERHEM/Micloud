package recordingmgr

import (
	"testing"
)

func TestListRecordings(t *testing.T) {
	recorders, err := ListRecordings(0, 10)
	if err != nil {
		t.Error(err)
	}
	if len(recorders) == 0 {
		t.Error("no recording")
	}
}

func TestGetRecording(t *testing.T) {
	url, err := GetRecordingFileUrl("43552880029221120")
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("url is empty")
	}
}

func TestDeleteRecording(t *testing.T) {
	err := DeleteRecording("43410534532923776")
	if err != nil {
		t.Error(err)
	}
}
