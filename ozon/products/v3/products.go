package v3

import (
	"net/http"

	"github.com/diphantxm/ozon-api-client/ozon/products/v3/info"
)

type SubRoutes struct {
	info *info.Info
}

func (c SubRoutes) Info() *info.Info {
	return c.info
}

func New(
	h *http.Client,
	uri string,
) *Products {
	return &Products{
		subRoutes: &SubRoutes{
			info: info.New(h, uri+"/info"),
		},
	}
}

type Products struct {
	subRoutes *SubRoutes
}

func (c Products) SubRoutes() *SubRoutes {
	return c.subRoutes
}
