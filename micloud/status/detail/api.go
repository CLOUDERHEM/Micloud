package detail

import (
	"fmt"
	"github.com/clouderhem/micloud/authorizer"
	"github.com/clouderhem/micloud/utility/validate"
	"net/http"
	"time"
)

const (
	detailApi = "https://i.mi.com/status/lite/alldetail?ts="
)

func GetAllDetail() (Detail, error) {
	req, err := http.NewRequest(http.MethodGet, detailApi+fmt.Sprint(time.Now().UnixMilli()), nil)
	if err != nil {
		return Detail{}, nil
	}

	body, r, err := authorizer.DoRequest(req)
	if err != nil {
		return Detail{}, nil
	}
	return validate.Validate[Detail](r, body)
}
