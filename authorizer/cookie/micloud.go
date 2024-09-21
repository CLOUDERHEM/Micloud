package cookie

import (
	"github.com/clouderhem/micloud/utility/parse"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	MicloudCookieFilepath = "/tmp/micloud/.micloud_cookie"
)

var (
	MicloudGlobalCookie = ""
	MicloudCookieLock   sync.Mutex
)

func GetValueFromMicloudCookie(name string) string {
	return parse.GetValueByKey(GetMicloudCookie(), name)
}

func GetMicloudCookie() string {
	if MicloudGlobalCookie != "" {
		return MicloudGlobalCookie
	}
	c := getCookieFromFile(MicloudCookieFilepath)
	MicloudCookieLock.Lock()
	MicloudGlobalCookie = c
	MicloudCookieLock.Unlock()
	return MicloudGlobalCookie
}

func SetMicloudCookie(s string) {
	s = strings.TrimSpace(s)
	MicloudCookieLock.Lock()
	MicloudGlobalCookie = s
	err := os.WriteFile(MicloudCookieFilepath, []byte(s), 0644)
	if err != nil {
		log.Printf("can not write cookie file: %v", err)
	}
	MicloudCookieLock.Unlock()
}
