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

func (c Info) Stocks(ctx context.Context, req *StocksRequest) (*StocksResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "StocksRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/stocks", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "StocksRequest.NewRequest")
	}

	return request.Send[StocksResponse](c.h, r, request.ContentTypeApplicationJson)
}
