package cookie

import (
	"io.github.clouderhem.micloud/consts"
	"io.github.clouderhem.micloud/utility/parse"
	"log"
	"os"
	"strings"
	"sync"
)

var GlobalCookie = ""

var lock sync.Mutex

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
