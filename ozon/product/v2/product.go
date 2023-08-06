package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/diphantxm/ozon-api-client/ozon/product/v2/info"

	"github.com/diphantxm/ozon-api-client/internal/request"
	"github.com/pkg/errors"
)

type SubRoutes struct {
	info *info.Info
}

func New(
	h *http.Client,
	uri string,
) *Product {
	return &Product{
		h:   h,
		uri: uri,

		subRoutes: &SubRoutes{
			info: info.New(h, uri+"/info"),
		},
	}
}

type Product struct {
	h   *http.Client
	uri string

	subRoutes *SubRoutes
}

func (c Product) SubRoutes() *SubRoutes {
	return c.subRoutes
}

func (c Product) Import(ctx context.Context, req *ImportRequest) (*ImportResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "ImportRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/import", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "ImportRequest.NewRequest")
	}

	return request.Send[ImportResponse](c.h, r, request.ContentTypeApplicationJson)
}

func (c Product) List(ctx context.Context, req *ListRequest) (*ListResponse, *http.Response, error) {
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

func (c Product) Info(ctx context.Context, req *InfoRequest) (*InfoResponse, *http.Response, error) {
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
