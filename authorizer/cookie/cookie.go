package cookie

import (
	"github.com/clouderhem/micloud/utility/parse"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	cookieFilepath = "/micloud/.micloud_cookie"
)

var (
	GlobalCookie = ""
	lock         sync.Mutex
)

func GetValue(name string) string {
	return parse.GetValueByKey(GetCookie(), name)
}

func GetCookie() string {
	if GlobalCookie != "" {
		return GlobalCookie
	}
	c := readCookieFile()
	lock.Lock()
	GlobalCookie = c
	lock.Unlock()
	return GlobalCookie
}

func SetCookie(s string) {
	s = strings.TrimSpace(s)
	lock.Lock()
	GlobalCookie = s
	err := os.WriteFile(cookieFilepath, []byte(s), 0644)
	if err != nil {
		log.Printf("can not create cookie file: %v", err)
	}
	lock.Unlock()
}

func readCookieFile() string {
	data, err := os.ReadFile(cookieFilepath)
	if os.IsNotExist(err) {
		err := os.WriteFile(cookieFilepath, []byte(""), 0644)
		if err != nil {
			log.Printf("can not create cookie file: %v", err)
			return ""
		}
	} else if err != nil {
		log.Printf("can not read cookie file: %v", err)
		return ""
	}
	return strings.TrimSpace(string(data))
}

func SetCookieFilepath(path string) {
	cookieFilepath = path
}
