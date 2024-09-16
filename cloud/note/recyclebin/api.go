package recyclebin

import (
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/cloud/note/note"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/validate"
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

	return validate.Validate[note.Notes](r, body)
}
