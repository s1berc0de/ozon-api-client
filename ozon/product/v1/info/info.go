package info

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/s1berc0de/ozon-api-client/ozon/product/v1/info/stocksbywarehouse"

	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
)

type SubRoutes struct {
	stocksByWarehouse *stocksbywarehouse.StocksByWarehouse
}

func (c SubRoutes) StocksByWarehouse() *stocksbywarehouse.StocksByWarehouse {
	return c.stocksByWarehouse
}

func New(
	h *http.Client,
	uri string,
) *Info {
	return &Info{
		h:   h,
		uri: uri,
		subRoutes: &SubRoutes{
			stocksByWarehouse: stocksbywarehouse.New(h, uri+"/stocks-by-warehouse"),
		},
	}
}

type Info struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c Info) Description(ctx context.Context, req *DescriptionRequest) (*DescriptionResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "DescriptionRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/description", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "DescriptionRequest.NewRequest")
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

func (c Info) Discounted(ctx context.Context, req *DiscountedRequest) (*DiscountedResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "DiscountedRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/discounted", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "DiscountedRequest.NewRequest")
	}

	return request.Send[DiscountedResponse](c.h, r, request.ContentTypeApplicationJson)
}
