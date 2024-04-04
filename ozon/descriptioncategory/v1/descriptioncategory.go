package v1

import (
	"net/http"

	"github.com/s1berc0de/ozon-api-client/ozon/descriptioncategory/v1/attribute"
	"github.com/s1berc0de/ozon-api-client/ozon/descriptioncategory/v1/tree"
)

type SubRoutes struct {
	attribute *attribute.Attribute
	tree      *tree.Tree
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
			tree:      tree.New(h, uri+"/tree"),
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
