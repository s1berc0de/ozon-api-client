package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/diphantxm/ozon-api-client/ozon/common"
	"github.com/pkg/errors"
	"net/http"
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

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/v3/category/attribute", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "TreeRequest.NewRequest")
	}

	return common.Send[AttributeResponse](c.h, r, common.ContentTypeApplicationJson)
}
