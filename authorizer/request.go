package authorizer

import (
	"io.github.clouderhem.micloud/utility/request"
	"net/http"
)

func DoRequest(req *http.Request) ([]byte, *http.Response, error) {
	postProcessReq(req)
	return request.DoRequest(req)
}

func postProcessReq(req *http.Request) {
	req.Header.Set("Cookie", GetCookie())
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Windows\"")
}
