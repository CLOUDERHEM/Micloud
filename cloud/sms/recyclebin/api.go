package recyclebin

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/cloud/sms/message"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
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

	if r.StatusCode != http.StatusOK {
		return message.Messages{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[message.Messages](body)
	if err != nil {
		return message.Messages{}, err
	}
	if resp.Code != consts.Success {
		return message.Messages{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
