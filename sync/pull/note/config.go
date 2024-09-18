package notesync

import (
	"errors"
	"net/url"

	"io.github.clouderhem.micloud/consts"
	mlog "io.github.clouderhem.micloud/utility/log"
)

const DirName = "note"

var (
	jsonFilepath = ""
	xlsxFilepath = ""
)

var (
	noteFailedFilepath = ""
	fileFailedFilepath = ""
)

var logFilepath = ""
var filesDirName = ""

var log mlog.Log

func init() {
	jsonFilepath = createPath("note.json")
	xlsxFilepath = createPath("note.xlsx")

	noteFailedFilepath = createPath("note_failed.json")
	fileFailedFilepath = createPath("file_failed.json")

	filesDirName = createPath("note")

	logger, err := mlog.NewFileLog(DirName, createPath(logFilepath))
	if err != nil {
		panic(errors.Join(errors.New("create note log file error"), err))
	}
	log = logger
}

func createPath(base string) string {
	filepath, err := url.JoinPath(consts.DatabaseDir, DirName, base)
	if err != nil {
		panic(errors.Join(errors.New("cannot build file path"), err))
	}
	return filepath
}
