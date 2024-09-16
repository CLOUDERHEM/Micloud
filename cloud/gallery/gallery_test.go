package gallery

import (
	"fmt"
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
