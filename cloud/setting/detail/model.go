package detail

type Detail struct {
	AutoRenewal                   bool        `json:"autoRenewal"`
	CurrentRecordIsLongTermRecord bool        `json:"currentRecordIsLongTermRecord"`
	Level                         string      `json:"level"`
	SettingType                   string      `json:"settingType"`
	TotalDetail                   TotalDetail `json:"totalDetail"`

	TotalQuota int64 `json:"totalQuota"`
	Used       int64 `json:"used"`

	UsedDetail UsedDetail `json:"usedDetail"`
}

type TotalDetail struct {
	BaseQuota               int64  `json:"baseQuota"`
	BonusSize               int64  `json:"bonusSize"`
	ExtendPackageSize       int64  `json:"extendPackageSize"`
	YearlyPackageExpireTime int64  `json:"yearlyPackageExpireTime"`
	YearlyPackageExpireSize int64  `json:"yearlyPackageExpireSize"`
	YearlyPackageExpireType string `json:"yearlyPackageExpireType"`
}

type UsedDetail struct {
	AppList      Record `json:"AppList"`
	GalleryImage Record `json:"GalleryImage"`
	Recorder     Record `json:"Recorder"`
}

type Record struct {
	Size int64  `json:"size"`
	Text string `json:"text"`
}
