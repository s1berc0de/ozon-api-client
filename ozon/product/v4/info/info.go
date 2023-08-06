package info

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
) *Info {
	return &Info{
		h:   h,
		uri: uri,
	}
}

type Info struct {
	h   *http.Client
	uri string
}

func (c Info) Limit(ctx context.Context, req *LimitRequest) (*LimitResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "LimitRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/limit", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "LimitRequest.NewRequest")
	}

	return request.Send[LimitResponse](c.h, r, request.ContentTypeApplicationJson)
}
