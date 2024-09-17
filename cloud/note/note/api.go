package note

import (
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/authorizer/cookie"
	"io.github.clouderhem.micloud/utility/validate"
	"net/http"
	"strconv"
	"time"
)

const (
	fullPageApi   = "https://i.mi.com/note/full/page/?limit=%v&ts=%v"
	noteApi       = "https://i.mi.com/note/note/%v/?ts=%v"
	deleteNoteApi = "https://i.mi.com/note/full/%v/delete"
)

func ListNotes(limit int) (Notes, error) {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(fullPageApi, limit, time.Now().UnixMilli()), nil)
	if err != nil {
		return Notes{}, err
	}
	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return Notes{}, err
	}
	return validate.Validate[Notes](r, body)
}

func GetNote(id string) (Note, error) {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(noteApi, id, time.Now().UnixMilli()), nil)
	if err != nil {
		return Note{}, err
	}
	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return Note{}, err
	}
	data, err := validate.Validate[EntryNote](r, body)
	if err != nil {
		return Note{}, err
	}
	return data.Entry, err
}

func DeleteNote(id, tag string, purge bool) error {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(deleteNoteApi, id), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded; charset=UTF-8")
	_ = req.ParseForm()
	req.Form.Add("tag", tag)
	req.Form.Add("purge", strconv.FormatBool(purge))
	req.Form.Add("serviceToken", cookie.GetValue("serviceToken"))

	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return err
	}

	_, err = validate.Validate[struct{}](r, body)
	if err != nil {
		return err
	}
	return nil
}
