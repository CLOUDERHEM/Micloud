package timeline

import (
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/validate"
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
	return validate.Validate[Timeline](r, body)
}
