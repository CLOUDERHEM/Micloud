package recorder

type Recorder struct {
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

type Recorders struct {
	List []Recorder `json:"list"`
}

type RecorderFile struct {
	Url string `json:"url"`
}
