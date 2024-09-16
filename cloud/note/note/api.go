package note

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/authorizer/cookie"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/response"
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
	if r.StatusCode != http.StatusOK {
		return Notes{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Notes](body)
	if err != nil {
		return Notes{}, err
	}
	if resp.Code != consts.Success {
		return Notes{}, errors.New(resp.Description)
	}
	return resp.Data, nil
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
	if r.StatusCode != http.StatusOK {
		return Note{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[EntryNote](body)
	if err != nil {
		return Note{}, err
	}
	if resp.Code != consts.Success {
		return Note{}, errors.New(resp.Description)
	}
	return resp.Data.Entry, nil
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
	if r.StatusCode != http.StatusOK {
		return errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[struct{}](body)
	if err != nil {
		return err
	}

	if resp.Code != consts.Success {
		return errors.New(resp.Description)
	}
	return nil
}
