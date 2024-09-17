package mlog

import (
	"log"
)

type StdLog struct {
	BaseLog
}

func NewStdLog(tag string) *StdLog {
	return &StdLog{
		BaseLog: BaseLog{
			tag:  tag,
			impl: log.Default(),
		},
	}
}
