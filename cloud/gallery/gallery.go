package gallery

import "io.github.clouderhem.micloud/cloud/gallery/album"

func ListAlbums(pageNum, pageSize int, shared bool) (album.Albums, error) {
	return album.ListAlbums(pageNum, pageSize, shared)
}
