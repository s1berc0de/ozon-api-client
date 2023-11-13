package v3

import (
	"net/http"

	"github.com/s1berc0de/ozon-api-client/ozon/product/v3/info"
)

type SubRoutes struct {
	info *info.Info
}

func (c SubRoutes) Import() *info.Info {
	return c.info
}

func New(
	h *http.Client,
	uri string,
) *Product {
	return &Product{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			info: info.New(h, uri+"/info"),
		},
	}
}

type Product struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c Product) SubRoutes() *SubRoutes {
	return c.subRoutes
}
