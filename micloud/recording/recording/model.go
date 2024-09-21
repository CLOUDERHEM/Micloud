package recording

type Recording struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ParentId string `json:"parent_id"`
	Size     int64  `json:"size"`
	Sha1     string `json:"sha1"`
	Ver      string `json:"ver"`

	CreateTime int64 `json:"create_time"`
	ModifyTime int64 `json:"modify_time"`
}

type Recordings struct {
	List []Recording `json:"list"`
}

type RecordingFile struct {
	Url string `json:"url"`
}
