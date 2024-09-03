package attribute

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/andmetoo/ozon-api-client/ozon/descriptioncategory/v1/attribute/values"
	"net/http"

	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
)

type SubRoutes struct {
	values *values.Values
}

func (c SubRoutes) Values() *values.Values {
	return c.values
}

func New(
	h *http.Client,
	uri string,
) *Attribute {
	return &Attribute{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			values: values.New(h, uri+"/values"),
		},
	}
}

type Attribute struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (a Attribute) SubRoutes() *SubRoutes {
	return a.subRoutes
}

func (a Attribute) Values(ctx context.Context, req *ValuesRequest) (*ValuesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ValuesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, a.uri+"/values", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ValuesRequest.NewRequest")
	}

	return request.Send[ValuesResponse](a.h, r, request.ContentTypeApplicationJson)
}
