package product

import (
	"net/http"

	v3 "github.com/s1berc0de/ozon-api-client/ozon/product/v3"

	v1 "github.com/s1berc0de/ozon-api-client/ozon/product/v1"
	v2 "github.com/s1berc0de/ozon-api-client/ozon/product/v2"
	v4 "github.com/s1berc0de/ozon-api-client/ozon/product/v4"
)

func New(
	h *http.Client,
	uri string,
) *Product {
	return &Product{
		v1: v1.New(h, uri+"/v1/product"),
		v2: v2.New(h, uri+"/v2/product"),
		v3: v3.New(h, uri+"/v3/product"),
		v4: v4.New(h, uri+"/v4/product"),
	}
}

type Product struct {
	v1 *v1.Product
	v2 *v2.Product
	v3 *v3.Product
	v4 *v4.Product
}

func (c *Product) V1() *v1.Product {
	return c.v1
}

func (c *Product) V2() *v2.Product {
	return c.v2
}

func (c *Product) V3() *v3.Product {
	return c.v3
}

func (c *Product) V4() *v4.Product {
	return c.v4
}
