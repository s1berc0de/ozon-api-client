package _import

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/andmetoo/ozon-api-client/internal/request"
	"github.com/pkg/errors"
)

func New(
	h *http.Client,
	uri string,
) *Import {
	return &Import{
		h:   h,
		uri: uri,
	}
}

type Import struct {
	h   *http.Client
	uri string
}

func (c Import) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/info", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.NewRequest")
	}

	return request.Send[InfoResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Import) Stocks(ctx context.Context, req *StocksRequest) (*StocksResponse, *http.Response, error) {
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

func (c Import) Prices(ctx context.Context, req *PricesRequest) (*PricesResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "PricesRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/prices", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "PricesRequest.NewRequest")
	}

	return request.Send[PricesResponse](c.h, r, request.ContentTypeApplicationJson)
}
