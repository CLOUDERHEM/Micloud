package cookie

import (
	"io.github.clouderhem.micloud/consts"
	"log"
	"os"
	"strings"
	"sync"
)

var lock sync.Mutex

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
	s = strings.TrimSpace(s)
	lock.Lock()
	err := os.WriteFile(consts.MicloudCookieFilePath, []byte(s), 0644)
	if err != nil {
		log.Printf("can not create cookie file: %v", err)
	}
	lock.Unlock()
}

func readCookieFile() string {
	data, err := os.ReadFile(consts.MicloudCookieFilePath)
	if os.IsNotExist(err) {
		err := os.WriteFile(consts.MicloudCookieFilePath, []byte(""), 0644)
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
