package parse

import (
	"fmt"
	"strings"
)

type KV struct {
	K string
	V string
}

func GetValueByKey(str, key string) string {
	values := GetKeyValues(str)
	for i := range values {
		if values[i].K == key {
			return values[i].V
		}
	}
	return ""
}

func GetKeyValues(str string) []KV {
	var result []KV
	kvs := strings.Split(str, ";")
	for i := range kvs {
		kv := strings.Split(kvs[i], "=")
		if len(kv) < 2 {
			continue
		}
		k := strings.TrimSpace(kv[0])
		v := strings.Join(kv[1:], "=")
		result = append(result, KV{k, v})
	}
	return result
}

// TidyKvs use the not empty v if s has same ks,
// use the last one if multi vs
func TidyKvs(s string) string {
	kvs := GetKeyValues(s)
	mp := make(map[string]string)
	for i := range kvs {
		if mp[kvs[i].K] == "" || mp[kvs[i].K] == "\"\"" {
			mp[kvs[i].K] = kvs[i].V
		}
	}
	var lines []string
	for k, v := range mp {
		lines = append(lines, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(lines, ";")
}
