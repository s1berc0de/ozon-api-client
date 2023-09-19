package v1

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
) *Warehouse {
	return &Warehouse{
		h:   h,
		uri: uri,
	}
}

type Warehouse struct {
	h   *http.Client
	uri string
}

func (c Warehouse) List(ctx context.Context, req *ListRequest) (*ListResponse, *http.Response, error) {
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
