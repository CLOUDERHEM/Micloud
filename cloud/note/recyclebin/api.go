package recyclebin

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/cloud/note/note"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
	"strconv"
	"time"
)

const (
	noteDeletedPageApi = "https://i.mi.com/note/deleted/page"
)

func ListDeletedNotes(syncTag string, limit int) (note.Notes, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"_dc", ts},
		{"limit", strconv.Itoa(limit)},
		{"syncTag", syncTag},
	}
	body, r, err := authorizer.DoRequest(request.NewGet(noteDeletedPageApi, q))
	if err != nil {
		return note.Notes{}, err
	}

	if r.StatusCode != http.StatusOK {
		return note.Notes{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[note.Notes](body)
	if err != nil {
		return note.Notes{}, err
	}
	if resp.Code != consts.Success {
		return note.Notes{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
