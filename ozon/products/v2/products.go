package v2

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
) *Products {
	return &Products{
		h:   h,
		uri: uri,
	}
}

type Products struct {
	h   *http.Client
	uri string
}

func (c Products) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, *http.Response, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, nil, errors.Wrap(err, "DeleteRequest.Marshal")
	}

	r, err := http.NewRequestWithContext(ctx, http.MethodPost, c.uri+"/delete", bytes.NewReader(b))
	if err != nil {
		return nil, nil, errors.Wrap(err, "DeleteRequest.NewRequest")
	}

	return request.Send[DeleteResponse](c.h, r, request.ContentTypeApplicationJson)
}
