package mlog

import "log"

type Log interface {
	LogI(v ...any)
	LogE(msg string, err error)
}

type BaseLog struct {
	tag  string
	impl *log.Logger
}

func (l *BaseLog) LogI(v ...any) {
	l.impl.Printf("[%v][I] %v", l.tag, v)
}

func (l *BaseLog) LogE(msg string, err error) {
	l.impl.Printf("[%v][E] %v, err: %v", l.tag, msg, err)
}
