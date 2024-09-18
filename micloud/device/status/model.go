package status

type Status struct {
	Devices []Device `json:"devices"`
	Locale  string   `json:"locale"`
}

type Device struct {
	CommandList           []string          `json:"commandList"`
	DevId                 string            `json:"devId"`
	DeviceType            string            `json:"deviceType"`
	Imei                  string            `json:"imei"`
	IsTZDevice            bool              `json:"isTZDevice"`
	LastLocationReceipt   LocationReceipt   `json:"lastLocationReceipt"`
	LastResponse          Response          `json:"lastResponse"`
	LocationReceiptList   []LocationReceipt `json:"locationReceiptList"`
	Model                 string            `json:"model"`
	Phone                 string            `json:"phone"`
	RegId                 string            `json:"regId"`
	ShowUnavailableNotice bool              `json:"showUnavailableNotice"`
	Snapshot              int               `json:"snapshot"`
	Status                string            `json:"status"`
	UpdateTime            int64             `json:"updateTime"`
	Version               string            `json:"version"`
}

type LocationReceipt struct {
	GpsInfo            GPSInfo   `json:"gpsInfo"`
	GpsInfoTransformed []GPSInfo `json:"gpsInfoTransformed"`
	InfoTime           int64     `json:"infoTime"`
	// fuck the shit, int | string
	MsgId             any    `json:"msgId"`
	Phone             string `json:"phone"`
	PowerLevel        int    `json:"powerLevel"`
	ResponseType      string `json:"responseType"`
	ServerReceiveTime int64  `json:"serverReceiveTime"`
}

type GPSInfo struct {
	Accuracy        float64 `json:"accuracy"`
	Address         string  `json:"address"`
	Area            string  `json:"area"`
	CoordinateType  string  `json:"coordinateType"`
	InChinaMainLand bool    `json:"inChinaMainLand"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	SourceType      string  `json:"sourceType"`
}

type Response struct {
	LastOperationTime int64  `json:"lastOperationTime"`
	LastResponseType  string `json:"lastResponseType"`
}
