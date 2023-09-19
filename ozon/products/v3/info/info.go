package info

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

func (c Info) Attributes(ctx context.Context, req *AttributesRequest) (*AttributesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "AttributesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/attributes", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "AttributesRequest.NewRequest")
	}

	return request.Send[AttributesResponse](c.h, r, request.ContentTypeApplicationJson)
}
