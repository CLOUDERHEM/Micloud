package contact

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
	initData = "https://i.mi.com/contacts/initdata"
)

func ListContacts(limit int) (Contacts, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"_dc", ts},
		{"syncTag", "0"},
		{"limit", strconv.Itoa(limit)},
		{"syncIgnoreTag", "0"},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(initData, q))
	if err != nil {
		return Contacts{}, err
	}

	if r.StatusCode != http.StatusOK {
		return Contacts{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Contacts](body)
	if err != nil {
		return Contacts{}, err
	}
	if resp.Code != consts.Success {
		return Contacts{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
