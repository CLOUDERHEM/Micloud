package mjson

import "github.com/tidwall/gjson"

func ListKeysAndValues(json []byte) (keys []string, values []string) {
	r := gjson.Parse(string(json))
	r.ForEach(func(key, value gjson.Result) bool {
		keys = append(keys, key.String())
		values = append(values, value.String())
		return true
	})
	return
}
