package recorder

import "io.github.clouderhem.micloud/cloud/recorder/recorders"

func ListRecorders(offset, limit int) ([]recorders.Recorder, error) {
	return recorders.ListRecorders(offset, limit)
}

func GetRecorderUrl(id string) (string, error) {
	return recorders.GetRecorderUrl(id)
}

func DeleteRecorder(id string) error {
	return recorders.DeleteRecorder(id)
}
