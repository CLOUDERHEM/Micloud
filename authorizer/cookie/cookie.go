package cookie

import (
	"log"
	"os"
	"strings"
	"sync"
)

var lock sync.Mutex

const (
	CookieFilePath = "/tmp/micloud/.micloud_cookie"
)

func GetCookieByName(name string) string {
	cookies := strings.Split(GetCookie(), ";")
	for i := range cookies {
		kv := strings.Split(cookies[i], "=")
		if len(kv) < 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.Join(kv[1:], "=")
		if k == name {
			return v
		}
	}
	return ""
}

func GetCookie() string {
	return readCookieFile()
}

func SetCookie(s string) {
	lock.Lock()
	err := os.WriteFile(CookieFilePath, []byte(s), 0644)
	if err != nil {
		log.Printf("can not create cookie file: %v", err)
	}
	lock.Unlock()
}

func readCookieFile() string {
	data, err := os.ReadFile(CookieFilePath)
	if os.IsNotExist(err) {
		err := os.WriteFile(CookieFilePath, []byte(""), 0644)
		if err != nil {
			log.Printf("can not create cookie file: %v", err)
			return ""
		}
	} else if err != nil {
		log.Printf("can not read cookie file: %v", err)
		return ""
	}
	return string(data)
}
