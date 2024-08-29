package products

import (
	"net/http"

	v1 "github.com/andmetoo/ozon-api-client/ozon/products/v1"

	v2 "github.com/andmetoo/ozon-api-client/ozon/products/v2"

	v3 "github.com/andmetoo/ozon-api-client/ozon/products/v3"
)

func New(
	h *http.Client,
	uri string,
) *Products {
	return &Products{
		v1: v1.New(h, uri+"/v1/products"),
		v2: v2.New(h, uri+"/v2/products"),
		v3: v3.New(h, uri+"/v3/products"),
	}
}

type Products struct {
	v1 *v1.Products
	v2 *v2.Products
	v3 *v3.Products
}

func (c *Products) V1() *v1.Products {
	return c.v1
}

func (c *Products) V2() *v2.Products {
	return c.v2
}

func (c *Products) V3() *v3.Products {
	return c.v3
}
