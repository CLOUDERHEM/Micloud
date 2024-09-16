package message

import (
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/validate"
	"strconv"
	"time"
)

const (
	fullThreadApi = "https://i.mi.com/sms/full/thread"
)

func ListMessages(syncTag, syncThreadTag string, limit int) (Messages, error) {
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

	body, r, err := authorizer.DoRequest(request.NewGet(fullThreadApi, q))
	if err != nil {
		return Messages{}, err
	}
	return validate.Validate[Messages](r, body)
}
