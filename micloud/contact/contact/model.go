package contact

type Content struct {
	DisplayName string `json:"displayName"`
	Id          string `json:"id"`
	Name        struct {
		Formatted string `json:"formatted"`
	} `json:"name"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers"`

	Starred bool   `json:"starred"`
	Status  string `json:"status"`
	Tag     string `json:"tag"`
}

type PhoneNumber struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Contacts struct {
	Content       map[string]ContentWrapper `json:"content"`
	LetterIndex   map[string][]string       `json:"letterIndex"`
	SyncIgnoreTag string                    `json:"syncIgnoreTag"`
	SyncTag       string                    `json:"syncTag"`
	LastPage      bool                      `json:"lastPage"`
}

type ContentWrapper struct {
	Content    Content `json:"content"`
	Pinyin     string  `json:"pinyin"`
	Type       string  `json:"type"`
	CreateTime int64   `json:"createTime"`
	UpdateTime int64   `json:"updateTime"`
}
