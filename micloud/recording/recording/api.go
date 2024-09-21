package recording

import (
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"net/http"
	"strconv"
	"time"
)

const (
	listRecordingsApi   = "https://i.mi.com/sfs/ns/recorder/dir/0/list"
	getRecordingFileApi = "https://i.mi.com/sfs/ns/recorder/file/%v/storage/geturl?ts=%v"
	deleteRecordApi     = "https://i.mi.com/sfs/ns/recorder/file/%v/delete?ts=%v"
)

func ListRecordings(offset, limit int) ([]Recording, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"offset", strconv.Itoa(offset)},
		{"limit", strconv.Itoa(limit)},
		{"ts", ts},
		{"_dc", ts},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(listRecordingsApi, q))
	if err != nil {
		return nil, err
	}

	data, err := validate.Validate[Recordings](r, body)
	if err != nil {
		return nil, err
	}
	return data.List, err
}

func GetRecordingFileUrl(id string) (string, error) {
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf(getRecordingFileApi, id, time.Now().UnixMilli()), nil)
	if err != nil {
		return "", err
	}

	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return "", err
	}

	data, err := validate.Validate[RecordingFile](r, body)
	if err != nil {
		return "", err
	}
	return data.Url, nil
}

func DeleteRecording(id string) error {
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf(deleteRecordApi, id, time.Now().UnixMilli()), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type",
		"application/x-www-form-urlencoded; charset=UTF-8")
	_ = req.ParseForm()
	req.Form.Add("permanent", "false")
	req.Form.Add("serviceToken", cookie.GetValueFromMicloudCookie("serviceToken"))

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
