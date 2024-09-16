package gallery

import (
	"fmt"
	"io.github.clouderhem.micloud/cloud/gallery/gallery"
	"testing"
)

func TestListAlbums(t *testing.T) {
	albums, err := ListAlbums(0, 10, false)
	if err != nil {
		t.Error(err)
	}
	if len(albums.Albums) == 0 {
		t.Error("no albums")
	}
	fmt.Println(albums)
}

func TestListGalleries(t *testing.T) {
	galleries, err := ListGalleries(gallery.GalleriesQuery{
		StartDate: 20240701,
		EndDate:   202409,
		PageNum:   0,
		PageSize:  1,
		AlbumId:   "1",
	})
	if err != nil {
		t.Error(err)
	}
	if len(galleries.Galleries) == 0 {
		t.Error("no galleries")
	}
	fmt.Println(galleries)
}

func TestGetTimeline(t *testing.T) {
	timeline, err := GetTimeline("1")
	if err != nil {
		t.Error(err)
	}
	if len(timeline.DayCount) == 0 {
		t.Error("no day count")
	}
	fmt.Println(timeline)
}

func TestDeleteGallery(t *testing.T) {
	err := DeleteGallery("1")
	if err != nil {
		t.Error(err)
	}
}

func TestGetGalleryFileUrl(t *testing.T) {
	url, err := GetGalleryFileUrl("1")
	if err != nil {
		t.Error(err)
	}
	if url == "" {
		t.Error("no url")
	}
	fmt.Println(url)
}
