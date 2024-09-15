package response

import (
	"testing"
)

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

func TestParse(t *testing.T) {
	s := "{\n    \"result\": \"ok\",\n    \"retriable\": false,\n    \"code\": 0,\n    \"data\": {\n        \"list\": [\n            {\n                \"devId\": \"devId\",\n                \"modelInfo\": {\n                    \"deviceType\": \"MI_PHONE\",\n                    \"modelName\": \"Xiaomi 12S Pro\",\n                    \"fullImageUrl\": \"https://cdn.cnbj1.fds.api.mi-img.com/device-model-img/L2S.png\",\n                    \"model\": \"Xiaomi 12S Pro\",\n                    \"deviceName\": \"小米手机\"\n                },\n                \"model\": \"2206122SC~Xiaomi 12S Pro\",\n                \"statusMicloud\": \"ok\"\n            }\n        ]\n    },\n    \"description\": \"成功\",\n    \"ts\": 123\n}"
	devices, err := Parse[Devices]([]byte(s))
	if err != nil {
		t.Error(err)
	}

	if len(devices.Data.List) != 1 {
		t.Error()
	}
	if devices.Data.List[0].DevId != "devId" {
		t.Error()
	}
}

func TestParseNoData(t *testing.T) {
	s := "{\"result\":\"ok\",\"retriable\":false,\"code\":-1,\"data\":\"\",\"description\":\"成功\",\"ts\":123}"
	devices, err := Parse[Devices]([]byte(s))
	if err != nil {
		t.Error(err)
	}

	if len(devices.Data.List) != 1 {
		t.Error()
	}
	if devices.Data.List[0].DevId != "devId" {
		t.Error()
	}
}
