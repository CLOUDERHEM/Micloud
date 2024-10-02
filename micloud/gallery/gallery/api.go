package gallery

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/utility/request"
	"github.com/clouderhem/micloud/utility/validate"
	"net/http"
	"strconv"
	"strings"
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

func GetGalleryStorageUrl(id string) (string, error) {
	ts := fmt.Sprintf("%d", time.Now().UnixMilli())
	q := []request.UrlQuery{
		{"ts", ts},
		{"id", id},
		{"callBack", fmt.Sprintf("dl_img_cb_%s_0", ts)},
	}

	body, r, err := authorizer.DoRequest(request.NewGet(storageApi, q))
	if err != nil {
		return "", err
	}

	data, err := validate.Validate[StorageFile](r, body)
	if err != nil {
		return "", err
	}
	return data.Url, nil
}

func GetGalleryFile(storageUrl string) (GalleryFile, error) {
	req, err := http.NewRequest(http.MethodGet, storageUrl, nil)
	if err != nil {
		return GalleryFile{}, err
	}
	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return GalleryFile{}, err
	}
	if r.StatusCode != http.StatusOK {
		return GalleryFile{}, errors.New(http.StatusText(r.StatusCode))
	}

	s := string(body)
	var file GalleryFile
	err = json.Unmarshal([]byte(s[strings.Index(s, "{"):len(s)-1]), &file)
	if err != nil {
		return GalleryFile{}, err
	}
	return file, nil
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

	_, err = validate.Validate[StorageFile](r, body)
	if err != nil {
		return err
	}
	return nil
}
