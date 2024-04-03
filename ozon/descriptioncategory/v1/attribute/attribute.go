package attribute

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
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

func (c Attribute) Attribute(ctx context.Context, req *AttributeRequest) (*AttributeResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "AttributeRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri, bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "AttributeRequest.NewRequest")
	}

	return request.Send[AttributeResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Attribute) Values(ctx context.Context, req *ValuesRequest) (*ValuesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ValuesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/values", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ValuesRequest.NewRequest")
	}

	return request.Send[ValuesResponse](c.h, r, request.ContentTypeApplicationJson)
}
