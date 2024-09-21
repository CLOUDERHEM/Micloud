package recordingmgr

import "github.com/clouderhem/micloud/micloud/recording/recording"

func ListRecordings(offset, limit int) ([]recording.Recording, error) {
	return recording.ListRecordings(offset, limit)
}

func GetRecordingFileUrl(id string) (string, error) {
	return recording.GetRecordingFileUrl(id)
}

func DeleteRecording(id string) error {
	return recording.DeleteRecording(id)
}
