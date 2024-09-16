package gallery

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
	galleriesApi     = "https://i.mi.com/gallery/user/galleries"
	storageApi       = "https://i.mi.com/gallery/storage"
	deleteGalleryApi = "https://i.mi.com/gallery/info/delete"
)

func ListGalleries(query GalleriesQuery) (Galleries, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"startDate", strconv.Itoa(query.StartDate)},
		{"endDate", strconv.Itoa(query.EndDate)},
		{"pageNum", strconv.Itoa(query.PageNum)},
		{"pageSize", strconv.Itoa(query.PageSize)},
		{"albumId", query.AlbumId},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(galleriesApi, q))
	if err != nil {
		return Galleries{}, err
	}

	if r.StatusCode != http.StatusOK {
		return Galleries{}, errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[Galleries](body)
	if err != nil {
		return Galleries{}, err
	}
	if resp.Code != consts.Success {
		return Galleries{}, errors.New(resp.Description)
	}
	return resp.Data, nil
}

func GetGalleryFileUrl(id string) (string, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"id", id},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(storageApi, q))
	if err != nil {
		return "", err
	}

	if r.StatusCode != http.StatusOK {
		return "", errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[GalleryFile](body)
	if err != nil {
		return "", err
	}
	if resp.Code != consts.Success {
		return "", errors.New(resp.Description)
	}
	return resp.Data.Url, nil
}

func DeleteGallery(id string) error {
	q := []request.UrlQuery{
		{"id", id},
		{"serviceToken", cookie.GetCookieByName("serviceToken")},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(deleteGalleryApi, q))
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return errors.New(http.StatusText(r.StatusCode))
	}

	resp, err := response.Parse[GalleryFile](body)
	if err != nil {
		return err
	}
	if resp.Code != consts.Success {
		return errors.New(resp.Description)
	}
	return nil
}
