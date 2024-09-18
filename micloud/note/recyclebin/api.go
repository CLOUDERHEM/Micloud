package recyclebin

import (
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/micloud/note/note"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
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
