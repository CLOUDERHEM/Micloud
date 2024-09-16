package parse

import "strings"

func GetValueByKey(str, key string) string {
	kvs := strings.Split(str, ";")
	for i := range kvs {
		kv := strings.Split(kvs[i], "=")
		if len(kv) < 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.Join(kv[1:], "=")
		if k == key {
			return v
		}
	}
	return ""
}
