package descriptioncategory

import (
	"net/http"

	v1 "github.com/andmetoo/ozon-api-client/ozon/descriptioncategory/v1"
)

func New(
	h *http.Client,
	uri string,
) *DescriptionCategory {
	return &DescriptionCategory{
		v1: v1.New(h, uri+"/v1/description-category"),
	}
}

type DescriptionCategory struct {
	v1 *v1.DescriptionCategory
}

func (c *DescriptionCategory) V1() *v1.DescriptionCategory {
	return c.v1
}
