package v1

import (
	"net/http"

	"github.com/s1berc0de/ozon-api-client/ozon/descriptioncategory/v1/attribute"
)

type SubRoutes struct {
	attribute *attribute.Attribute
}

func (c SubRoutes) Attribute() *attribute.Attribute {
	return c.attribute
}

func New(
	h *http.Client,
	uri string,
) *DescriptionCategory {
	return &DescriptionCategory{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			attribute: attribute.New(h, uri+"/attribute"),
		},
	}
}

type DescriptionCategory struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c DescriptionCategory) SubRoutes() *SubRoutes {
	return c.subRoutes
}
