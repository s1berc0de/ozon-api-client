package attribute

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/diphantxm/ozon-api-client/internal/ozon/request"
	"github.com/pkg/errors"
	"net/http"
)

func New(
	h *http.Client,
	uri string,
) *Attribute {
	return &Attribute{
		h:   h,
		uri: uri,
	}
}

type Attribute struct {
	h   *http.Client
	uri string
}

func (c Attribute) Values(ctx context.Context, req *ValuesRequest) (*ValuesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ValuesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/v2/category/attribute/values", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ValuesRequest.NewRequest")
	}

	return request.Send[ValuesResponse](c.h, r, request.ContentTypeApplicationJson)
}
