package gallerymgr

import (
	"errors"
	"github.com/clouderhem/micloud/micloud/gallery/album"
	"github.com/clouderhem/micloud/micloud/gallery/gallery"
	"github.com/clouderhem/micloud/micloud/gallery/timeline"
)

func ListAlbums(pageNum, pageSize int, shared bool) (album.Albums, error) {
	return album.ListAlbums(pageNum, pageSize, shared)
}

func ListGalleries(query gallery.GalleriesQuery) (gallery.Galleries, error) {
	return gallery.ListGalleries(query)
}

func GetGalleryFileUrl(id string) (string, error) {
	url, err := gallery.GetGalleryFileUrl(id)
	if err != nil {
		return "", err
	}
	if url == "" {
		return "", errors.New("cannot get gallery file url")
	}
	return url, err
}

func GetTimeline(albumId string) (timeline.Timeline, error) {
	return timeline.GetTimeline(albumId)
}

func DeleteGallery(id string) error {
	return gallery.DeleteGallery(id)
}
