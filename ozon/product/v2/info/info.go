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

func (c Info) List(ctx context.Context, req *ListRequest) (*ListResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ListRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/list", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ListRequest.NewRequest")
	}

	return request.Send[ListResponse](c.h, r, request.ContentTypeApplicationJson)
}
