package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
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

func (c DescriptionCategory) Tree(ctx context.Context, req *TreeRequest) (*TreeResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/tree", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.NewRequest")
	}

	return request.Send[TreeResponse](c.h, r, request.ContentTypeApplicationJson)
}
