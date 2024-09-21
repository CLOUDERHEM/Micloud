package gallery

import (
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
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
	return validate.Validate[Galleries](r, body)
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

	data, err := validate.Validate[GalleryFile](r, body)
	if err != nil {
		return "", err
	}
	return data.Url, nil
}

func DeleteGallery(id string) error {
	q := []request.UrlQuery{
		{"id", id},
		{"serviceToken", cookie.GetValueFromMicloudCookie("serviceToken")},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(deleteGalleryApi, q))
	if err != nil {
		return err
	}

	_, err = validate.Validate[GalleryFile](r, body)
	if err != nil {
		return err
	}
	return nil
}
