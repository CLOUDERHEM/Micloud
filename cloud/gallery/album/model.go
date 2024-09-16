package album

type Album struct {
	AlbumId        string `json:"albumId"`
	LastUpdateTime int64  `json:"lastUpdateTime"`
	MediaCount     int    `json:"mediaCount"`
	Name           string `json:"name"`
	UserId         int64  `json:"userId"`

	Thumbnails []Thumbnail `json:"thumbnails"`
}

type Thumbnail struct {
	Url         string `json:"url"`
	Orientation int    `json:"orientation"`
}

type Albums struct {
	Albums     []Album `json:"albums"`
	IndexHash  int     `json:"indexHash"`
	IsLastPage bool    `json:"isLastPage"`
}
