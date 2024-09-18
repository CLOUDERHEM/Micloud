package smsmgr

import (
	"github.com/clouderhem/micloud/micloud/sms/message"
	"github.com/clouderhem/micloud/micloud/sms/recyclebin"
)

func ListMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	return message.ListMessages(syncTag, syncThreadTag, limit)
}

func ListDeletedMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	return recyclebin.ListDeletedMessages(syncTag, syncThreadTag, limit)
}
