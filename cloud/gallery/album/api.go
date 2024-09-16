package album

import (
	"errors"
	"fmt"
	"io.github.clouderhem.micloud/authorizer"
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/request"
	"io.github.clouderhem.micloud/utility/response"
	"net/http"
	"strconv"
	"time"
)

const (
	listAlbumApi = "https://i.mi.com/gallery/user/album/list"
)

func ListAlbums(pageNum, pageSize int, shared bool) (Albums, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"_dc", ts},
		{"pageNum", strconv.Itoa(pageNum)},
		{"pageSize", strconv.Itoa(pageSize)},
		{"isShared", strconv.FormatBool(shared)},
		{"numOfThumbnails", "1"},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(listAlbumApi, q))
	if err != nil {
		return Albums{}, err
	}

	if r.StatusCode != http.StatusOK {
		return Albums{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Albums](body)
	if err != nil {
		return Albums{}, err
	}
	if resp.Code != consts.Success {
		return Albums{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}
