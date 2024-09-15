package recorders

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/authorizer/cookie"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
	"strconv"
	"time"
)

const (
	listRecorderApi  = "https://i.mi.com/sfs/ns/recorder/dir/0/list"
	getRecordFileApi = "https://i.mi.com/sfs/ns/recorder/file/%v/storage/geturl?ts=%v"
	deleteRecordApi  = "https://i.mi.com/sfs/ns/recorder/file/%v/delete?ts=%v"
)

func ListRecorders(offset, limit int) ([]Recorder, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"offset", strconv.Itoa(offset)},
		{"limit", strconv.Itoa(limit)},
		{"ts", ts},
		{"_dc", ts},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(listRecorderApi, q))
	if err != nil {
		return nil, err
	}
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Recorders](body)
	if err != nil {
		return nil, err
	}
	if resp.Code != consts.Success {
		return nil, errors.New(resp.Description)
	}
	return resp.Data.List, nil
}

func GetRecorderUrl(id string) (string, error) {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(getRecordFileApi, id, time.Now().UnixMilli()), nil)
	if err != nil {
		return "", err
	}

	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[RecorderFile](body)
	if err != nil {
		return "", err
	}
	if resp.Code != consts.Success {
		return "", errors.New(resp.Description)
	}
	return resp.Data.Url, nil
}

func DeleteRecorder(id string) error {
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf(deleteRecordApi, id, time.Now().UnixMilli()), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded; charset=UTF-8")
	_ = req.ParseForm()
	req.Form.Add("permanent", "false")
	req.Form.Add("serviceToken", cookie.GetCookieByName("serviceToken"))

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
