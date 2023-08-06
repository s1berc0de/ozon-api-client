package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/diphantxm/ozon-api-client/internal/request"
	"github.com/pkg/errors"
)

func New(
	h *http.Client,
	uri string,
) *Category {
	return &Category{
		h:   h,
		uri: uri,
	}
}

type Category struct {
	h   *http.Client
	uri string
}

func (c Category) Attribute(ctx context.Context, req *AttributeRequest) (*AttributeResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/attribute", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.NewRequest")
	}

	return request.Send[AttributeResponse](c.h, r, request.ContentTypeApplicationJson)
}
