package category

import (
	v2 "github.com/diphantxm/ozon-api-client/ozon/category/v2"
	v3 "github.com/diphantxm/ozon-api-client/ozon/category/v3"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Category {
	return &Category{
		v2: v2.New(h, uri),
		v3: v3.New(h, uri),
	}
}

type Category struct {
	v2 *v2.Category
	v3 *v3.Category
}

func (c Category) V2() *v2.Category {
	return c.v2
}

func (c Category) V3() *v3.Category {
	return c.v3
}
