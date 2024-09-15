package response

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"log"
)

type Response[T any] struct {
	Code        int    `json:"code"`
	Data        T      `json:"data"`
	Result      string `json:"result"`
	Description string `json:"description"`
}

func Parse[T any](body []byte) (Response[T], error) {
	parse := gjson.Parse(string(body))
	r := Response[T]{
		Code:        int(parse.Get("code").Int()),
		Result:      parse.Get("result").String(),
		Description: parse.Get("description").String(),
	}

	data := parse.Get("data")
	if data.Raw != "" {
		var t T
		err := json.Unmarshal([]byte(data.Raw), &t)
		if err != nil {
			log.Printf("cannot unmarshal r data, data: %v, err: %v", data.Raw, err)
		} else {
			r.Data = t
		}
	}

	return r, nil
}
