package recorder

import (
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/authorizer/cookie"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/validate"
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

	data, err := validate.Validate[Recorders](r, body)
	if err != nil {
		return nil, err
	}
	return data.List, err
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

	data, err := validate.Validate[RecorderFile](r, body)
	if err != nil {
		return "", err
	}
	return data.Url, nil
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
