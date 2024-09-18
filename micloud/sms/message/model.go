package message

type Message struct {
	Folder         any    `json:"folder"`
	Id             string `json:"id"`
	LastUpdateTime int64  `json:"lastUpdateTime"`
	LocalTime      int64  `json:"localTime"`
	Recipients     string `json:"recipients"`
	Snippet        string `json:"snippet"`
	Tag            string `json:"tag"`
	ThreadId       string `json:"threadId"`
	Total          int    `json:"total"`
	Unread         int    `json:"unread"`
}

type MsgEntry struct {
	Entry     Message `json:"entry"`
	Operation string  `json:"operation"`
}

type Messages struct {
	Entries   []MsgEntry `json:"entries"`
	Watermark Watermark  `json:"watermark"`
}

type Watermark struct {
	SyncTag       string `json:"syncTag"`
	SyncThreadTag string `json:"syncThreadTag"`
}
