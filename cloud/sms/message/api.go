package message

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
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

	if r.StatusCode != http.StatusOK {
		return Messages{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Messages](body)
	if err != nil {
		return Messages{}, err
	}
	if resp.Code != consts.Success {
		return Messages{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
