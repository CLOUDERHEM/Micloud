package authorizer

import (
	"github.com/clouderhem/micloud/authorizer/cookie"
	"github.com/clouderhem/micloud/config"
	"github.com/clouderhem/micloud/micloud/status/setting"
	"github.com/clouderhem/micloud/utility/request"
	"log"
	"net/http"
)

func DoRequest(req *http.Request) ([]byte, *http.Response, error) {
	var body []byte
	var resp *http.Response
	var err error
	for i := 0; i < int(config.RetryTimes); i++ {
		postProcessReq(req)
		body, resp, err = request.DoRequest(req)
		if err != nil {
			continue
		}
		if resp.StatusCode == http.StatusUnauthorized {
			c, err := setting.Renewal()
			if err != nil || c == "" {
				log.Print("cannot renewal cookie, err: ", err)
				return body, resp, err
			} else {
				cookie.SetCookie(c)
			}
			continue
		}
		break
	}

	return body, resp, err
}

func postProcessReq(req *http.Request) {
	req.Header.Set("Cookie", cookie.GetCookie())
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
}
