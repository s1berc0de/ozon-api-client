package warehouse

import (
	"net/http"

	v1 "github.com/diphantxm/ozon-api-client/ozon/warehouse/v1"
)

func New(
	h *http.Client,
	uri string,
) *Warehouse {
	return &Warehouse{
		v1: v1.New(h, uri+"/v1/warehouse"),
	}
}

type Warehouse struct {
	v1 *v1.Warehouse
}
