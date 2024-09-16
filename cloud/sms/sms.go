package sms

import (
	"io.github.clouderhem.micloud/cloud/sms/message"
	"io.github.clouderhem.micloud/cloud/sms/recyclebin"
)

func ListMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	return message.ListMessages(syncTag, syncThreadTag, limit)
}

func ListDeletedMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	return recyclebin.ListDeletedMessages(syncTag, syncThreadTag, limit)
}
