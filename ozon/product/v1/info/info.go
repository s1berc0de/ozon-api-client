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

func (c Info) Description(ctx context.Context, req *DescriptionRequest) (*DescriptionResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/description", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "InfoRequest.NewRequest")
	}

	return request.Send[DescriptionResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Info) Subscription(ctx context.Context, req *SubscriptionRequest) (*SubscriptionResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "SubscriptionRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/subscription", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "SubscriptionRequest.NewRequest")
	}

	return request.Send[SubscriptionResponse](c.h, r, request.ContentTypeApplicationJson)
}
