package cookie

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getCookieFromFile(cookiePath string) string {
	data, err := os.ReadFile(cookiePath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(cookiePath), os.ModePerm)
		if err != nil {
			log.Fatalf("cannot mkdir for cookie path, path: %v, err: %v", cookiePath, err)
		}
		err = os.WriteFile(cookiePath, []byte(""), 0644)
		if err != nil {
			log.Fatalf("can not create cookie file, path: %v, err: %v", cookiePath, err)
			return ""
		}
	} else if err != nil {
		log.Fatalf("can not read cookie file, path: %v, err: %v", cookiePath, err)
		return ""
	}
	return strings.TrimSpace(string(data))
}
