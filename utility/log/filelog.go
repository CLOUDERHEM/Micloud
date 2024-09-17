package mlog

import (
	"log"
	"os"
	"path/filepath"
)

type FileLog struct {
	LogPath string
	BaseLog
}

func NewFileLog(tag, path string) (*FileLog, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &FileLog{
		LogPath: path,
		BaseLog: BaseLog{
			tag:  tag,
			impl: log.New(file, "", log.LstdFlags),
		},
	}, nil
}
