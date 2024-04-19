package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/s1berc0de/ozon-api-client/internal/request"
	"net/http"

	"github.com/s1berc0de/ozon-api-client/ozon/product/v3/info"
)

type SubRoutes struct {
	info *info.Info
}

func (c SubRoutes) Info() *info.Info {
	return c.info
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
		return nil, nil, errors.Wrap(err, "Import.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/import", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "Import.NewRequest")
	}

	return request.Send[ImportResponse](c.h, r, request.ContentTypeApplicationJson)
}
