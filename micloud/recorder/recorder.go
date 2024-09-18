package recordermgr

import "github.com/clouderhem/micloud/micloud/recorder/recorder"

func ListRecorders(offset, limit int) ([]recorder.Recorder, error) {
	return recorder.ListRecorders(offset, limit)
}

func GetRecorderUrl(id string) (string, error) {
	return recorder.GetRecorderUrl(id)
}

func DeleteRecorder(id string) error {
	return recorder.DeleteRecorder(id)
}
