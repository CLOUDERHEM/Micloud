package note

type Note struct {
	AlertDate  int64 `json:"alertDate"`
	AlertTag   int   `json:"alertTag"`
	ColorId    int   `json:"colorId"`
	CreateDate int64 `json:"createDate"`
	// fuck the shit, it can be int or string
	FolderId   any     `json:"folderId"`
	Id         string  `json:"id"`
	ModifyDate int64   `json:"modifyDate"`
	Setting    Setting `json:"setting"`
	Snippet    string  `json:"snippet"`
	Status     string  `json:"status"`
	Subject    string  `json:"subject"`
	Tag        string  `json:"tag"`
	Type       string  `json:"type"`

	Content   string `json:"content"`
	ExtraInfo string `json:"extraInfo"`
}

type Folder struct {
	AlertDate  int64 `json:"alertDate"`
	AlertTag   int   `json:"alertTag"`
	ColorId    int   `json:"colorId"`
	CreateDate int64 `json:"createDate"`
	DeleteTime int64 `json:"deleteTime"`
	// fucking the shit, it can be int or string
	FolderId   any     `json:"folderId"`
	Id         string  `json:"id"`
	ModifyDate int64   `json:"modifyDate"`
	Setting    Setting `json:"setting"`
	Snippet    string  `json:"snippet"`
	Status     string  `json:"status"`
	Subject    string  `json:"subject"`
	Tag        string  `json:"tag"`
	Type       string  `json:"type"`
}

type Setting struct {
	ThemeId    int   `json:"themeId"`
	Version    int   `json:"version"`
	StickyTime int64 `json:"stickyTime"`
}

type Notes struct {
	Entries  []Note   `json:"entries"`
	Folders  []Folder `json:"folders"`
	LastPage bool     `json:"lastPage"`
	SyncTag  string   `json:"syncTag"`
}

type EntryNote struct {
	Entry Note `json:"entry"`
}
