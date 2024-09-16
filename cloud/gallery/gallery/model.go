package gallery

type ThumbnailInfo struct {
	Data  string `json:"data"`
	IsUrl bool   `json:"isUrl"`
}

type Gallery struct {
	BigThumbnailInfo ThumbnailInfo `json:"bigThumbnailInfo"`
	CreateTime       int64         `json:"createTime"`
	DateModified     int64         `json:"dateModified"`
	DateToken        int64         `json:"dateToken"`
	Description      string        `json:"description"`
	FileName         string        `json:"fileName"`
	GeoInfo          GeoInfo       `json:"geoInfo"`
	GroupId          any           `json:"groupId"`
	Id               string        `json:"id"`
	IsFavorite       bool          `json:"isFavorite"`
	IsFrontCamera    bool          `json:"isFrontCamera"`
	IsUbiImage       bool          `json:"isUbiImage"`
	MimeType         string        `json:"mimeType"`
	Sha1             string        `json:"sha1"`
	Size             int64         `json:"size"`
	SortTime         int64         `json:"sortTime"`
	Status           string        `json:"status"`
	Tag              string        `json:"tag"`
	ThumbnailInfo    ThumbnailInfo `json:"thumbnailInfo"`
	Title            string        `json:"title"`
	Type             string        `json:"type"`
}

type ExifInfo struct {
	FNumber             string `json:"FNumber"`
	GPSAltitude         string `json:"GPSAltitude"`
	GPSAltitudeRef      any    `json:"GPSAltitudeRef"`
	GPSDateStamp        string `json:"GPSDateStamp"`
	GPSLatitude         string `json:"GPSLatitude"`
	GPSLatitudeRef      any    `json:"GPSLatitudeRef"`
	GPSLongitude        string `json:"GPSLongitude"`
	GPSLongitudeRef     any    `json:"GPSLongitudeRef"`
	GPSProcessingMethod string `json:"GPSProcessingMethod"`
	GPSTimeStamp        string `json:"GPSTimeStamp"`
	ISOSpeedRatings     string `json:"ISOSpeedRatings"`
	DateTime            string `json:"dateTime"`
	ExposureTime        string `json:"exposureTime"`
	Flash               int    `json:"flash"`
	FocalLength         string `json:"focalLength"`
	ImageLength         int    `json:"imageLength"`
	ImageWidth          int    `json:"imageWidth"`
	Make                string `json:"make"`
	Model               string `json:"model"`
	Orientation         int    `json:"orientation"`
}

type Galleries struct {
	Galleries  []Gallery `json:"galleries"`
	IndexHash  int64     `json:"indexHash"`
	IsLastPage bool      `json:"isLastPage"`
}

type GeoInfo struct {
	Address     Address   `json:"address"`
	AddressList []Address `json:"addressList"`
	Gps         string    `json:"gps"`
	IsAccurate  bool      `json:"isAccurate"`
}

type Address struct {
	AddressLines    []string `json:"addressLines"`
	Admin           string   `json:"admin"`
	CountryName     string   `json:"countryName"`
	Locale          string   `json:"locale"`
	Locality        string   `json:"locality"`
	SubLocality     string   `json:"subLocality"`
	SubThoroughfare string   `json:"subThoroughfare"`
	Thoroughfare    string   `json:"thoroughfare"`
}

type GalleriesQuery struct {
	StartDate int `json:"startDate"`
	EndDate   int `json:"endDate"`
	PageNum   int `json:"pageNum"`
	PageSize  int `json:"pageSize"`
	// camera 1
	AlbumId string `json:"albumId"`
}

type GalleryFile struct {
	Url string `json:"url"`
}
