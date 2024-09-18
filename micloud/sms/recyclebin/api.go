package recyclebin

import (
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/micloud/sms/message"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"strconv"
	"time"
)

const (
	deletedThread = "https://i.mi.com/sms/deleted/thread"
)

func ListDeletedMessages(syncTag, syncThreadTag string, limit int) (message.Messages, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"_dc", ts},
		{"limit", strconv.Itoa(limit)},
		{"syncTag", syncTag},
		{"syncThreadTag", syncThreadTag},
		{"readMode", "older"},
		{"withPhoneCall", "false"},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(deletedThread, q))
	if err != nil {
		return message.Messages{}, err
	}
	return validate.Validate[message.Messages](r, body)
}
