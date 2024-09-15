package device

type Device struct {
	DevId         string    `json:"devId"`
	Model         string    `json:"model"`
	ModelInfo     ModelInfo `json:"modelInfo"`
	StatusMicloud string    `json:"statusMicloud"`
}

type ModelInfo struct {
	DeviceName   string `json:"deviceName"`
	DeviceType   string `json:"deviceType"`
	FullImageUrl string `json:"fullImageUrl"`
	Model        string `json:"model"`
	ModelName    string `json:"modelName"`
}

type Devices struct {
	List []Device `json:"list"`
}
