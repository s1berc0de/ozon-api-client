package attributes

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
) *Attributes {
	return &Attributes{
		h:   h,
		uri: uri,
	}
}

type Attributes struct {
	h   *http.Client
	uri string
}

func (c Attributes) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "UpdateRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/update", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "UpdateRequest.NewRequest")
	}

	return request.Send[UpdateResponse](c.h, r, request.ContentTypeApplicationJson)
}
