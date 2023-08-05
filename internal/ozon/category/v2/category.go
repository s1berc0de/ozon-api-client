package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/diphantxm/ozon-api-client/internal/ozon/category/v2/attribute"
	"github.com/diphantxm/ozon-api-client/internal/ozon/request"
	"github.com/pkg/errors"
	"net/http"
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
) *Category {
	return &Category{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			attribute: attribute.New(h, uri),
		},
	}
}

type Category struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c Category) SubRoutes() *SubRoutes {
	return c.subRoutes
}

func (c Category) Tree(ctx context.Context, req *TreeRequest) (*TreeResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/v2/category/tree", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.NewRequest")
	}

	return request.Send[TreeResponse](c.h, r, request.ContentTypeApplicationJson)
}
