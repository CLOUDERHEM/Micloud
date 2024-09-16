package timeline

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
	"time"
)

const (
	timelineApi = "https://i.mi.com/gallery/user/timeline"
)

func GetTimeline(albumId string) (Timeline, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"albumId", albumId},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(timelineApi, q))
	if err != nil {
		return Timeline{}, err
	}

	if r.StatusCode != http.StatusOK {
		return Timeline{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Timeline](body)
	if err != nil {
		return Timeline{}, err
	}
	if resp.Code != consts.Success {
		return Timeline{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
