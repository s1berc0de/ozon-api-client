package category

import (
	"net/http"

	v2 "github.com/andmetoo/ozon-api-client/ozon/category/v2"
	v3 "github.com/andmetoo/ozon-api-client/ozon/category/v3"
)

func New(
	h *http.Client,
	uri string,
) *Category {
	return &Category{
		v2: v2.New(h, uri+"/v2/category"),
		v3: v3.New(h, uri+"/v3/category"),
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
