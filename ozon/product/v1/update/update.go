package update

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
) *Update {
	return &Update{
		h:   h,
		uri: uri,
	}
}

type Update struct {
	h   *http.Client
	uri string
}

func (c Update) OfferID(ctx context.Context, req *OfferIDRequest) (*OfferIDResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "OfferIDRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/offer-id", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "OfferIDRequest.NewRequest")
	}

	return request.Send[OfferIDResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Update) Discount(ctx context.Context, req *DiscountRequest) (*DiscountResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "DiscountRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/discount", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "DiscountRequest.NewRequest")
	}

	return request.Send[DiscountResponse](c.h, r, request.ContentTypeApplicationJson)
}
