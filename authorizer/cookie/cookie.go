package cookie

import (
	"log"
	"os"
	"strings"
)

func getCookieFromFile(cookiePath string) string {
	data, err := os.ReadFile(cookiePath)
	if os.IsNotExist(err) {
		err := os.WriteFile(cookiePath, []byte(""), 0644)
		if err != nil {
			log.Printf("can not create cookie file, path: %v, err: %v", cookiePath, err)
			return ""
		}
	} else if err != nil {
		log.Printf("can not read cookie file, path: %v, err: %v", cookiePath, err)
		return ""
	}
	return strings.TrimSpace(string(data))
}
