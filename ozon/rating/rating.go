package rating

import (
	"net/http"

	v1 "github.com/diphantxm/ozon-api-client/ozon/rating/v1"
)

func New(
	h *http.Client,
	uri string,
) *Rating {
	return &Rating{
		v1: v1.New(h, uri+"/v1/rating"),
	}
}

type Rating struct {
	v1 *v1.Rating
}
